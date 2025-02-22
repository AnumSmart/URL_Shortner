package stat

import (
	"server/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(dataBase *db.Db) *StatRepository {
	return &StatRepository{
		Db: dataBase,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	//если статистики по сслыке нет, то создаём
	//если есть, то увеличиваем на 1
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.DB.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.DB.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repo.DB.Save(&stat)
	}
}

func (repo *StatRepository) GetStats(by string, from, to time.Time) []GetStatResponce {
	var stats []GetStatResponce
	var selectQuerry string
	switch by {
	case GroupByDay:
		selectQuerry = "to_char(date, `YYYY-MM-DD`) as period, sum(clicks)"
	case GroupByMonth:
		selectQuerry = "to_char(date, `YYYY-MM`) as period, sum(clicks)"
	}
	repo.DB.Table("ststs").
		Select(selectQuerry).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats
}

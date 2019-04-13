package models

import (
	"log"
	"math"
	"os"
	"time"

	"goadmin/libs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ICDIV   = "idlecrafting"
	RPDIV   = "report"
	SMDIV   = "smash"
	BZDIV   = "banzi"
	GNDIV   = "gunner"
	QSDIV   = "qise"
	QS99DIV = "qise99"
	TGDIV   = "triangle"
	EMDIV   = "electro"
	CADIV   = "card2048"
	PTDIV   = "pintu"
)

// InitOrm ...
func InitOrm() {
	initOrm("default")
	initOrm(ICDIV)
	initOrm(RPDIV)
	initOrm(SMDIV)
	initOrm(BZDIV)
	initOrm(GNDIV)
	initOrm(QSDIV)
	initOrm(QS99DIV)
	initOrm(TGDIV)
	initOrm(EMDIV)
	initOrm(CADIV)
	initOrm(PTDIV)
	sql2log("mysql")
}

func sql2log(driverName string) {
	debug, _ := beego.AppConfig.Bool(driverName + ".debug")
	logfile := beego.AppConfig.String(driverName + ".logfile")
	orm.Debug = debug
	f, err := os.Create(logfile)
	if err != nil {
		log.Panicf("sql2log err: %v\n", err)
	}
	orm.DebugLog = orm.NewLog(f)
}

func initOrm(aliasName string) {
	driverName := beego.AppConfig.String(aliasName + ".driverName")
	dataSource := beego.AppConfig.String(aliasName + ".dataSource")
	maxIdle, _ := beego.AppConfig.Int(aliasName + ".maxIdle")
	maxConn, _ := beego.AppConfig.Int(aliasName + ".maxConn")
	location := beego.AppConfig.String(aliasName + ".location")
	prefix := beego.AppConfig.String(aliasName + ".prefix")

	orm.RegisterDriver(driverName, orm.DRMySQL)
	switch aliasName {
	case ICDIV:
		orm.RegisterModelWithPrefix(prefix,
			new(IcBaseData),
			new(IcBasePet),
			new(IcBaseMaterials),
			new(IcBaseEffect),
			new(IcBaseSkill),
			new(IcBaseAchiement),
			new(IcBaseTask),
			new(IcBaseBuyTimes),
			new(IcBaseChest),
			new(GameTask),
			new(GameAchiement),
		)
	case EMDIV:
		orm.RegisterModelWithPrefix(prefix,
			new(EmBaseData),
			new(EmBaseFu),
			new(GameFuDai),
			new(GameFuGate),
			new(GameFuDay),
		)
	case PTDIV:
		orm.RegisterModelWithPrefix(prefix,
			new(PtBaseData),
			new(PtBaseFu),
		)
	case TGDIV:
		orm.RegisterModelWithPrefix(prefix,
			new(TgBaseData),
		)
	case CADIV:
		orm.RegisterModelWithPrefix(prefix,
			new(CaBaseData),
		)
	case GNDIV:
		orm.RegisterModelWithPrefix(prefix,
			new(GnBaseData),
		)
	case BZDIV:
		orm.RegisterModelWithPrefix(prefix,
			new(BzBaseData),
			new(BzBaseAction),
		)
	case QSDIV:
		orm.RegisterModelWithPrefix(prefix,
			new(QsBaseData),
		)
	case QS99DIV:
		orm.RegisterModelWithPrefix(prefix,
			new(Qs99BaseData),
		)
	case SMDIV:
		orm.RegisterModelWithPrefix(prefix,
			new(SmBaseData),
		)
	case RPDIV:
		orm.RegisterModelWithPrefix(prefix,
			new(StatActive),
			new(StatPlaytime),
			new(StatStart),
			new(StatSendClick),
			new(StatProgress),
			new(StatScene),
			new(StatGate),
			new(StatPet),
			new(StatUseSkill),
			new(StatGateGrow),
			new(StatGateStop),
			new(StatGateData),
			new(StatOnline),
			new(StatShare),
			new(StatGateGold),
			new(StatDraw),
			new(StatWall),
			new(StatFast),
			new(StatSign),
			new(StatAstrology),
			//
			new(StatGateStop2),
			new(StatGate2),
			new(StatStart2),
			new(StatGateScore),
			new(StatGateRetimes),
			//
			new(LogLogin),
			new(LogRegister),
			new(LogSendClick),
			new(LogProgress),
			new(LogSceneAct),
			new(LogPlaytime),
			new(LogTip),
			new(LogWall),
			new(LogGateGrow),
			new(LogGold),
			new(LogGate),
			new(StatGateLoss),
			new(StatGateWeapon),
			new(StatPetLevel),
			new(StatGateGoldDis),
			new(StatGateGoldDep),
			new(StatWeaponLev),
			new(StatProp),
		)
	case "default":
		orm.RegisterModelWithPrefix(prefix,
			new(SystemUser),
			new(SystemModule),
			new(UserRole),
			new(UserPermission),
			new(SystemConfig),
			new(Project),
			new(SystemChat),
			new(ApiDoc),
			new(Server),
		)
	default:
	}
	orm.DefaultTimeLoc, _ = time.LoadLocation(location)
	orm.RegisterDataBase(aliasName, driverName, dataSource, maxIdle, maxConn)
}

// newOrm create new orm, switch default
func newOrm() orm.Ormer {
	return orm.NewOrm()
}

// switch new orm by driver
func newOrmBy(driverName string) orm.Ormer {
	o := orm.NewOrm()
	err := o.Using(driverName)
	if err != nil {
		beego.Error("newOrmBy err: ", err)
	}
	return o
}

// get prefix by driver
func prefixBy(driverName string) string {
	return beego.AppConfig.String(driverName + ".prefix")
}

// get prefix by projectID
func prefixByID(projectID int) (prefix string) {
	switch projectID {
	case 10:
		prefix = prefixBy(SMDIV)
	case 11:
		prefix = prefixBy(ICDIV)
	case 16:
		prefix = prefixBy(BZDIV)
	case 17:
		prefix = prefixBy(QSDIV)
	case 18:
		prefix = prefixBy(QS99DIV)
	case 19:
		prefix = prefixBy(GNDIV)
	case 20:
		prefix = prefixBy(TGDIV)
	case 21:
		prefix = prefixBy(EMDIV)
	case 22:
		prefix = prefixBy(CADIV)
	case 23:
		prefix = prefixBy(PTDIV)
	default:
		prefix = prefixBy("default")
	}
	return
}

// get orm and prefix by projectID
func getOrmBy(projectID int) (o orm.Ormer, prefix string) {
	switch projectID {
	case 10:
		o = newOrmBy(SMDIV)
		prefix = prefixBy(SMDIV)
	case 11:
		o = newOrmBy(ICDIV)
		prefix = prefixBy(ICDIV)
	case 16:
		o = newOrmBy(BZDIV)
		prefix = prefixBy(BZDIV)
	case 17:
		o = newOrmBy(QSDIV)
		prefix = prefixBy(QSDIV)
	case 18:
		o = newOrmBy(QS99DIV)
		prefix = prefixBy(QS99DIV)
	case 19:
		o = newOrmBy(GNDIV)
		prefix = prefixBy(GNDIV)
	case 20:
		o = newOrmBy(TGDIV)
		prefix = prefixBy(TGDIV)
	case 21:
		o = newOrmBy(EMDIV)
		prefix = prefixBy(EMDIV)
	case 22:
		o = newOrmBy(CADIV)
		prefix = prefixBy(CADIV)
	case 23:
		o = newOrmBy(PTDIV)
		prefix = prefixBy(PTDIV)
	default:
		o = newOrm()
		prefix = prefixBy("default")
	}
	return
}

// queryPage page > 0
func queryPage(qs orm.QuerySeter, page, limit int, exprs ...string) orm.QuerySeter {
	if len(exprs) != 0 {
		qs = qs.OrderBy(exprs...)
	} else if len(exprs) == 0 {
		qs = qs.OrderBy("-id")
	}
	return qs.Offset(((page - 1) * limit)).Limit(limit)
}

// queryCreatedAt2 指定日期按天查询
func queryCreatedAt2(qs orm.QuerySeter, datetime string) orm.QuerySeter {
	startTime, endTime := libs.Datetime2Str(datetime)
	qs = queryCreatedAt(qs, startTime, endTime)
	return qs
}

// queryCreatedAt 指定查询时间范围
func queryCreatedAt(qs orm.QuerySeter, startTime, endTime string) orm.QuerySeter {
	if startTime != "" && endTime != "" {
		st, err := libs.Str2LocalTime(startTime)
		if err != nil {
			beego.Error("queryCreatedAt err: ", err)
			return qs
		}
		et, err := libs.Str2LocalTime(endTime)
		if err != nil {
			beego.Error("queryCreatedAt err: ", err)
			return qs
		}
		qs = qs.Filter("created_at__gte", st)
		qs = qs.Filter("created_at__lt", et)
	} else if startTime != "" {
		st, err := libs.Str2LocalTime(startTime)
		if err != nil {
			beego.Error("queryCreatedAt err: ", err)
			return qs
		}
		qs = qs.Filter("created_at__gte", st)
	} else if endTime != "" {
		et, err := libs.Str2LocalTime(endTime)
		if err != nil {
			beego.Error("queryCreatedAt err: ", err)
			return qs
		}
		qs = qs.Filter("created_at__lt", et)
	}
	return qs
}

// set pagination
func setPag(page int, count int64) Pagination {
	return Pagination{
		CurrentPage: int32(page),
		TotalCounts: int32(math.Ceil(float64(count) / 20)),
	}
}

// set pagination
func setPag2(page, limit int, count int64) Pagination {
	if limit <= 0 {
		limit = 20
	}
	return Pagination{
		CurrentPage: int32(page),
		TotalCounts: int32(math.Ceil(float64(count) / float64(limit))),
	}
}

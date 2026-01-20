package sqlite

import (
	"context"
	"time"

	"github.com/SmurfsAtWork/lilpapa/app"
	"github.com/SmurfsAtWork/lilpapa/app/models"
	"github.com/SmurfsAtWork/lilpapa/evy"
	"github.com/SmurfsAtWork/lilpapa/nanoid"
	"gorm.io/gorm"
)

var _ app.Repository = &Repository{}

type Repository struct {
	client *gorm.DB
}

func New() (*Repository, error) {
	db, err := dbConnector()
	if err != nil {
		return nil, err
	}

	return &Repository{
		client: db,
	}, nil
}

// --------------------------------
// App Repository
// --------------------------------

func (r *Repository) CreateUser(user models.User) (models.User, error) {
	err := tryWrapDbError(
		gorm.G[models.User](r.client).
			Create(context.Background(), &user),
	)
	if _, ok := err.(*ErrRecordExists); ok {
		return models.User{}, &app.ErrExists{
			ResourceName: "user",
		}
	}
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) GetUser(id uint) (models.User, error) {
	user, err := gorm.G[models.User](r.client).
		Where("id = ?", id).
		First(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return models.User{}, &app.ErrNotFound{
			ResourceName: "user",
		}
	}
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserByUsername(username string) (models.User, error) {
	user, err := gorm.G[models.User](r.client).
		Where("username = ?", username).
		First(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return models.User{}, &app.ErrNotFound{
			ResourceName: "user",
		}
	}
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) CreateSmurf(smurf models.Smurf) (models.Smurf, error) {
	smurf.NanoId = nanoid.New()

	for range 25 {
		_, err := r.GetSmurfByNanoId(smurf.NanoId)
		if err != nil {
			break
		}
		smurf.NanoId = nanoid.New()
	}

	err := tryWrapDbError(
		gorm.G[models.Smurf](r.client).
			Create(context.Background(), &smurf),
	)
	if _, ok := err.(*ErrRecordExists); ok {
		return models.Smurf{}, &app.ErrExists{
			ResourceName: "smurf",
		}
	}
	if err != nil {
		return models.Smurf{}, err
	}

	return smurf, nil
}

func (r *Repository) GetSmurf(id uint) (models.Smurf, error) {
	smurf, err := gorm.G[models.Smurf](r.client).
		Where("id = ?", id).
		First(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return models.Smurf{}, &app.ErrNotFound{
			ResourceName: "smurf",
		}
	}
	if err != nil {
		return models.Smurf{}, err
	}

	return smurf, nil
}

func (r *Repository) GetSmurfByNanoId(nanoId string) (models.Smurf, error) {
	smurf, err := gorm.G[models.Smurf](r.client).
		Where("nano_id = ?", nanoId).
		First(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return models.Smurf{}, &app.ErrNotFound{
			ResourceName: "smurf",
		}
	}
	if err != nil {
		return models.Smurf{}, err
	}

	return smurf, nil
}

func (r *Repository) UpdateSmurfPassword(id uint, newPassword string) error {
	_, err := gorm.G[models.Smurf](r.client).
		Where("id = ?", id).
		Update(context.Background(), "password", newPassword)
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return &app.ErrNotFound{
			ResourceName: "smurf",
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteSmurf(id uint) error {
	_, err := gorm.G[models.Smurf](r.client).
		Where("id = ?", id).
		Delete(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return &app.ErrNotFound{
			ResourceName: "smurf",
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpsertSmurfConfig(smurfId uint, smurfConfig models.SmurfConfig) (models.SmurfConfig, error) {
	config, err := gorm.G[models.SmurfConfig](r.client).
		Where("smurf_id = ?", smurfId).
		First(context.Background())
	if err != nil {
		err := gorm.G[models.SmurfConfig](r.client).Create(context.Background(), &smurfConfig)
		if err != nil {
			return models.SmurfConfig{}, err
		}

		return smurfConfig, nil
	}

	_, err = gorm.G[models.SmurfConfig](r.client).
		Where("id = ?", config.Id).
		Updates(context.Background(), smurfConfig)
	if err != nil {
		return models.SmurfConfig{}, err
	}

	return smurfConfig, nil
}

func (r *Repository) GetSmurfConfig(smurfId uint) (models.SmurfConfig, error) {
	config, err := gorm.G[models.SmurfConfig](r.client).
		Where("smurf_id = ?", smurfId).
		First(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return models.SmurfConfig{}, &app.ErrNotFound{
			ResourceName: "smurf_config",
		}
	}
	if err != nil {
		return models.SmurfConfig{}, err
	}

	return config, nil
}

func (r *Repository) CreateSmurfCommand(smurfCommand models.SmurfCommand) (models.SmurfCommand, error) {
	err := tryWrapDbError(
		gorm.G[models.SmurfCommand](r.client).
			Create(context.Background(), &smurfCommand),
	)
	if _, ok := err.(*ErrRecordExists); ok {
		return models.SmurfCommand{}, &app.ErrExists{
			ResourceName: "smurf_command",
		}
	}
	if err != nil {
		return models.SmurfCommand{}, err
	}

	return smurfCommand, nil
}

func (r *Repository) GetSmurfCommand(smurfId uint) (models.SmurfCommand, error) {
	command, err := gorm.G[models.SmurfCommand](r.client).
		Where("smurf_id = ?", smurfId).
		First(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return models.SmurfCommand{}, &app.ErrNotFound{
			ResourceName: "smurf_command",
		}
	}
	if err != nil {
		return models.SmurfCommand{}, err
	}

	return command, nil
}

func (r *Repository) GetSmurfCommands(smurfId uint) ([]models.SmurfCommand, error) {
	commands, err := gorm.G[models.SmurfCommand](r.client).
		Where("smurf_id = ?", smurfId).
		Find(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return nil, &app.ErrNotFound{
			ResourceName: "smurf_command",
		}
	}
	if err != nil {
		return nil, err
	}

	return commands, nil
}

func (r *Repository) DeleteSmurfCommand(smurfId uint, commandId uint) error {
	_, err := gorm.G[models.SmurfCommand](r.client).
		Where("smurf_id = ? AND id = ?", smurfId, commandId).
		Delete(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return &app.ErrNotFound{
			ResourceName: "smurf_command",
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteSmurfCommands(smurfId uint) error {
	_, err := gorm.G[models.SmurfCommand](r.client).
		Where("smurf_id = ?", smurfId).
		Delete(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return &app.ErrNotFound{
			ResourceName: "smurf_command",
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateSmurfLog(smurfLog models.SmurfLog) (models.SmurfLog, error) {
	err := tryWrapDbError(
		gorm.G[models.SmurfLog](r.client).
			Create(context.Background(), &smurfLog),
	)
	if _, ok := err.(*ErrRecordExists); ok {
		return models.SmurfLog{}, &app.ErrExists{
			ResourceName: "smurf_log",
		}
	}
	if err != nil {
		return models.SmurfLog{}, err
	}

	return smurfLog, nil
}

func (r *Repository) GetSmurfLogs(smurfId uint, since time.Time) ([]models.SmurfLog, error) {
	logs, err := gorm.G[models.SmurfLog](r.client).
		Where("smurf_id = ? AND created_at > ?", smurfId, since).
		Find(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return nil, &app.ErrNotFound{
			ResourceName: "smurf_log",
		}
	}
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *Repository) CreateSmurfStat(smurfStat models.SmurfStatus) (models.SmurfStatus, error) {
	err := tryWrapDbError(
		gorm.G[models.SmurfStatus](r.client).
			Create(context.Background(), &smurfStat),
	)
	if _, ok := err.(*ErrRecordExists); ok {
		return models.SmurfStatus{}, &app.ErrExists{
			ResourceName: "smurf_stat",
		}
	}
	if err != nil {
		return models.SmurfStatus{}, err
	}

	return smurfStat, nil
}

func (r *Repository) GetSmurfStats(smurfId uint, since time.Time) ([]models.SmurfStatus, error) {
	stats, err := gorm.G[models.SmurfStatus](r.client).
		Where("smurf_id = ? AND created_at > ?", smurfId, since).
		Find(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return nil, &app.ErrNotFound{
			ResourceName: "smurf_stat",
		}
	}
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *Repository) CreateProgram(program models.Program) (models.Program, error) {
	err := tryWrapDbError(
		gorm.G[models.Program](r.client).
			Create(context.Background(), &program),
	)
	if _, ok := err.(*ErrRecordExists); ok {
		return models.Program{}, &app.ErrExists{
			ResourceName: "program",
		}
	}
	if err != nil {
		return models.Program{}, err
	}

	return program, nil
}

func (r *Repository) CreateScript(script models.Script) (models.Script, error) {
	err := tryWrapDbError(
		gorm.G[models.Script](r.client).
			Create(context.Background(), &script),
	)
	if _, ok := err.(*ErrRecordExists); ok {
		return models.Script{}, &app.ErrExists{
			ResourceName: "script",
		}
	}
	if err != nil {
		return models.Script{}, err
	}

	return script, nil
}

func (r *Repository) CreateRunnable(runnable models.Runnable) (models.Runnable, error) {
	err := tryWrapDbError(
		gorm.G[models.Runnable](r.client).
			Create(context.Background(), &runnable),
	)
	if _, ok := err.(*ErrRecordExists); ok {
		return models.Runnable{}, &app.ErrExists{
			ResourceName: "runnable",
		}
	}
	if err != nil {
		return models.Runnable{}, err
	}

	return runnable, nil
}

// --------------------------------
// Evy Repository
// --------------------------------

func (r *Repository) CreateEvent(e evy.EventPayload) error {
	err := tryWrapDbError(
		gorm.G[evy.EventPayload](r.client).
			Create(context.Background(), &e),
	)
	if _, ok := err.(*ErrRecordExists); ok {
		return &app.ErrExists{
			ResourceName: "event",
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetEventsBatch(size int32) ([]evy.EventPayload, error) {
	events, err := gorm.G[evy.EventPayload](r.client).
		Limit(int(size)).
		Find(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return nil, &app.ErrNotFound{
			ResourceName: "event",
		}
	}
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return nil, &app.ErrNotFound{
			ResourceName: "event",
		}
	}

	return events, nil
}

func (r *Repository) DeleteEvent(id uint) error {
	_, err := gorm.G[evy.EventPayload](r.client).
		Where("id = ?", id).
		Delete(context.Background())
	err = tryWrapDbError(err)
	if _, ok := err.(*ErrRecordNotFound); ok {
		return &app.ErrNotFound{
			ResourceName: "event",
		}
	}
	if err != nil {
		return err
	}

	return nil
}

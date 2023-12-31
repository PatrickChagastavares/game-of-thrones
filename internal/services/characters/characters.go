//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package characters

import (
	"context"

	"github.com/PatrickChagastavares/game-of-thrones/internal/entities"
	"github.com/PatrickChagastavares/game-of-thrones/internal/repositories"
	"github.com/PatrickChagastavares/game-of-thrones/pkg/logger"
	"github.com/PatrickChagastavares/game-of-thrones/pkg/tracer"
)

type (
	IService interface {
		Create(ctx context.Context, newCharacter entities.CharacterRequest) (id string, err error)
		Find(ctx context.Context) (characters []entities.Character, err error)
		FindByID(ctx context.Context, id string) (character entities.Character, err error)
		Update(ctx context.Context, updateCharacter entities.CharacterRequest) (character entities.Character, err error)
		Delete(ctx context.Context, id string) (err error)
	}

	services struct {
		repositories *repositories.Container
		log          logger.Logger
	}
)

func New(repo *repositories.Container, log logger.Logger) IService {
	return &services{repositories: repo, log: log}
}

func (srv *services) Create(ctx context.Context, newCharacter entities.CharacterRequest) (id string, err error) {
	ctx, span := tracer.Span(ctx, "services.characters.create")
	defer span.End()

	newCharacter.PreSave(ctx)

	err = srv.repositories.Database.Character.Create(ctx, newCharacter)
	if err != nil {
		srv.log.ErrorContext(ctx, "character.Service.database.Create", err, ", playload: ", newCharacter)
		return id, err
	}

	return newCharacter.ID, nil
}

func (srv *services) Find(ctx context.Context) (characters []entities.Character, err error) {
	ctx, span := tracer.Span(ctx, "services.characters.find")
	defer span.End()

	characters, err = srv.repositories.Database.Character.Find(ctx)
	if err != nil {
		srv.log.ErrorContext(ctx, "character.Service.database.Find", err)
		return nil, ErrFind
	}

	return characters, nil
}

func (srv *services) FindByID(ctx context.Context, id string) (character entities.Character, err error) {
	ctx, span := tracer.Span(ctx, "services.characters.findbyid")
	defer span.End()

	character, err = srv.repositories.Database.Character.FindByID(ctx, id)
	if err != nil {
		srv.log.ErrorContext(ctx, "character.Service.database.FindByID", err)
		return character, ErrCharacterNotFound
	}

	return character, nil
}

func (srv *services) Update(ctx context.Context, updateCharacter entities.CharacterRequest) (character entities.Character, err error) {
	ctx, span := tracer.Span(ctx, "services.characters.update")
	defer span.End()

	character, err = srv.FindByID(ctx, updateCharacter.ID)
	if err != nil {
		return
	}

	character.PreUpdate(ctx, updateCharacter)

	err = srv.repositories.Database.Character.Update(ctx, &character)
	if err != nil {
		srv.log.ErrorContext(ctx, "character.Service.database.Update", err)
		return character, err
	}

	return character, nil
}

func (srv *services) Delete(ctx context.Context, id string) (err error) {
	ctx, span := tracer.Span(ctx, "services.characters.delete")
	defer span.End()

	_, err = srv.FindByID(ctx, id)
	if err != nil {
		return
	}

	err = srv.repositories.Database.Character.Delete(ctx, id)
	if err != nil {
		srv.log.ErrorContext(ctx, "character.Service.database.Delete", err)
		return err
	}

	if err := srv.repositories.Database.House.RemoveLord(ctx, id); err != nil {
		srv.log.ErrorContext(ctx, "character.Service.database.House.RemoveLord", err)
		return err
	}

	return nil
}

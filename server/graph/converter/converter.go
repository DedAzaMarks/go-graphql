package converter

import (
	"github.com/DedAzaMarks/hw07-graphql/server/graph/model"
	"github.com/DedAzaMarks/hw07-graphql/server/repository"
)

func GameToGame(game repository.Game) *model.Game {
	return &model.Game{
		ID:         game.Id.String(),
		Scoreboard: ScoreboardToScoreboard(game.Scoreboard),
		Comments:   game.Comments,
	}
}

func ScoreboardToScoreboard(scoreboard []repository.Player) []*model.Player {
	res := make([]*model.Player, 0, len(scoreboard))
	for _, s := range scoreboard {
		res = append(res, PlayerToPlayer(s))
	}
	return res
}

func PlayerToPlayer(player repository.Player) *model.Player {
	return &model.Player{
		Username:  player.Username,
		Role:      RoleToRole(player.Role),
		IsAlive:   player.IsAlive,
		DaysAlive: player.DaysAlive,
	}
}

func RoleToRole(role repository.Role) model.Role {
	switch role {
	case repository.Civilian:
		return model.RoleCivilian
	case repository.Mafia:
		return model.RoleMafia
	default:
		return "<nil>"
	}
}

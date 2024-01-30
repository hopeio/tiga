package binding

import (
	"github.com/gofiber/fiber/v3"
	stringsi "github.com/hopeio/tiga/utils/strings"
)

type Ctx fiber.Ctx

func (c *Ctx) Peek(key string) ([]string, bool) {
	ctx := (fiber.Ctx)(c)
	v := stringsi.BytesToString(ctx.Request().URI().QueryArgs().Peek(key))
	if v != "" {
		return []string{v}, true
	}
	v = ctx.Params(key)
	return []string{v}, v != ""
}

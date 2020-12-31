package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}

package action

import (
	"fmt"
	"errors"
	"os/exec"
	"os"
	"log"
	ulog "github.com/Unknwon/log"
)

func ApportGoto(args []string) error {
	ulog.Debug("goto:",args)

	if len(args) ==0{
		return errors.New("No where to go?")
	}

	cfg := DecodeCfg()

	for _,toItem := range cfg.To{
		var checkOk = true
		for j,aliasName := range toItem.Alias{
			if args[j] != aliasName{
				checkOk = false
				break
			}
		}

		if checkOk{
			doRedirect(toItem)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("No alias %v defined.", args))
}

func doRedirect(to ToItem){
	ulog.Debug("redirect:",to)

	finalArguments := []string{"/k","start","cmd","/k","cd","/d",to.Target}

	cmd := exec.Command("cmd", finalArguments...)
	cmd.Stdout = os.Stdout
	if err:=cmd.Run(); err!=nil{
		log.Fatal(err)
	}
}
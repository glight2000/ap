package action

import (
	"fmt"
	"github.com/Unknwon/log"
	"errors"
	"os/exec"
	"regexp"
	"os"
)

func ApportRun(args []string) error {
	log.Debug("run",args)
	cfg := DecodeCfg()

	if(len(args) == 0){
		return errors.New("Nothing to run?")
	}

	shortcut := args[0]

	for _,customize := range cfg.Customize{
		if shortcut == customize.Shortcut{
			log.Debug("to run command", shortcut)
			for _,target := range customize.Targets{
				argumentFilter := target.Argument_filter
				if len(args)>1{
					//TODO: only first argument is checked by filters, this should be improved
					//若shortcut后边有参数，参数与filter正则对比，找到合适的target
					match,_ := regexp.MatchString(argumentFilter, args[1])
					if match{
						doOsExcute(target, args[1:])
						return nil
					}
				}else{
					//查找filter为空的target
					if(argumentFilter ==""){
						doOsExcute(target, []string{})
						return nil
					}
				}
			}
		}
	}

	return errors.New(fmt.Sprintf("[%s] is not defined as a shortcut in customize or no filter accepted the arguments.", shortcut))
}

func doOsExcute(target CustomizeTarget, followArguments []string){
	log.Debug("doOsExcute")
	log.Debug("target:",target)
	log.Debug("followArguments:", followArguments)

	var finalArguments []string
	if target.IsArgumentsInherit {
		finalArguments = append(target.Ext_arguments, followArguments...)
	}else{
		finalArguments = target.Ext_arguments
	}

	//TODO: some application can not run (like lnk shortcut) or will blocks cmd window after started

	log.Debug("finalArguments:", finalArguments)

	cmd := exec.Command(target.App, finalArguments...)
	cmd.Stdout = os.Stdout
	cmd.Start()
}


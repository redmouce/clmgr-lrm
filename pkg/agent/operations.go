package agent

import (
	"time"

	"github.com/google/logger"
	. "myproj.com/clmgr-lrm/pkg/common"
	"os"
)

const (
	def_start_timeout   = time.Duration(20 * time.Second)
	def_stop_timeout    = time.Duration(20 * time.Second)
	def_monitor_timeout = time.Duration(20 * time.Second)
)

const (
	svn_operation = "CLM_OPERATION"
)

func (ag *agent) Start() error {
	logger.Infof("Starting agent %s", ag.Name())
	e := NewExecutor()
	if isIn, act := GetFromSliceF(ToInterface(ag.Config.Actions), at_start, func(x interface{}) interface{} {
		return x.(Action).Name
	}); isIn != -1 {
		e.SetTimeout(act.(Action).Timeout)
		os.Setenv(svn_operation+ag.Name(), "start")
		e.SetOp([]string{ag.scriptPath})
		_, err := e.Exec()
		if err != nil {
			logger.Errorf("The error accrued during start: %s", err.Error())
			if act.(Action).OnFail != of_ignore {
				logger.Info("On fail not ignore is STUBBED")
			} else {
				logger.Error("Default action is ignore")
			}
			return err
		}
	} else {
		// default way to start a resource is to perform it using systemd
		e.SetTimeout(def_start_timeout)
		e.SetOp([]string{"systemctl", "start", ag.Name()})
		_, err := e.Exec()
		if err != nil {
			logger.Errorf("The error accrued during start: %s", err.Error())
			logger.Error("Default action is ignore")
			return err
		}
	}
	return nil
}

func (ag *agent) Stop() error {
	logger.Infof("Stopping agent %s", ag.Name())
	e := NewExecutor()
	if isIn, act := GetFromSlice(ToInterface(ag.Config.Actions), at_stop); isIn != -1 {
		e.SetTimeout(act.(Action).Timeout)
		os.Setenv(svn_operation+ag.Name(), "stop")
		e.SetOp([]string{ag.scriptPath})
		_, err := e.Exec()
		if err != nil {
			logger.Errorf("The error accrued during stop: %s", err.Error())
			if act.(Action).OnFail != of_ignore {
				logger.Info("On fail not ignore is STUBBED")
			} else {
				logger.Error("Default action is ignore")
			}
			return err
		}
	} else {
		// default way to start a resource is to perform it using systemd
		e.SetTimeout(def_stop_timeout)
		e.SetOp([]string{"systemctl", "stop", ag.Name()})
		_, err := e.Exec()
		if err != nil {
			logger.Errorf("The error accrued during start: %s", err.Error())
			logger.Error("Default action is ignore")
			return err
		}
	}
	return nil
}

func (ag *agent) Monitor() interface{} {
	logger.Infof("Monitor on %s", ag.Name())
	e := NewExecutor()
	if isIn, act := GetFromSlice(ToInterface(ag.Config.Actions), at_monitor); isIn != -1 {
		e.SetTimeout(act.(Action).Timeout)
		os.Setenv(svn_operation+ag.Name(), "monitor")
		e.SetOp([]string{ag.scriptPath})
		res, err := e.Exec()
		if err != nil {
			logger.Errorf("The error accrued during monitor: %s", err.Error())
			if act.(Action).OnFail != of_ignore {
				logger.Info("On fail not ignore is STUBBED")
			} else {
				logger.Error("Default action is ignore")
			}
			return err
		}
		return res
	} else {
		// default way to start a resource is to perform it using systemd
		e.SetTimeout(def_monitor_timeout)
		e.SetOp([]string{"systemctl", "status", ag.Name()})
		res, err := e.Exec()
		if err != nil {
			logger.Errorf("The error accrued during start: %s", err.Error())
			logger.Error("Default action is ignore")
			return err
		}
		return res
	}
	return nil
}

func (ag *agent) Notify() error {
	logger.Infof("Notify %s", ag.Name())
	e := NewExecutor()
	if isIn, act := GetFromSlice(ToInterface(ag.Config.Actions), at_notify); isIn != -1 {
		e.SetTimeout(act.(Action).Timeout)
		os.Setenv(svn_operation+ag.Name(), "notify")
		e.SetOp([]string{ag.scriptPath})
		_, err := e.Exec()
		if err != nil {
			logger.Errorf("The error accrued during notify: %s", err.Error())
			if act.(Action).OnFail != of_ignore {
				logger.Info("On fail not ignore is STUBBED")
			} else {
				logger.Error("Default action is ignore")
			}
			return err
		}
	}
	logger.Info("Notify operation is not provided")
	return nil
}

func (ag *agent) Reload() error {
	logger.Infof("Resource %s Reload op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Promote() error {
	logger.Infof("Resource %s Promote op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Demote() error {
	logger.Infof("Resource %s Demote op is stubbed", ag.Name())
	return nil
}

func (ag *agent) MetaData() interface{} {
	logger.Infof("meta-data on %s", ag.Name())
	e := NewExecutor()
	if isIn, act := GetFromSlice(ToInterface(ag.Config.Actions), at_monitor); isIn != -1 {
		e.SetTimeout(act.(Action).Timeout)
		os.Setenv(svn_operation+ag.Name(), "meta-data")
		e.SetOp([]string{ag.scriptPath})
		res, err := e.Exec()
		if err != nil {
			logger.Errorf("The error accrued during meta-data: %s", err.Error())
			if act.(Action).OnFail != of_ignore {
				logger.Info("On fail not ignore is STUBBED")
			} else {
				logger.Error("Default action is ignore")
			}
			return err
		}
		return res
	}
	logger.Info("Metha-data is not provided")
	return nil
}

package main

const (
	Success = iota
	ErrParseConfiguration
	ErrPreHookRuntime
	ErrPostHookRuntime
	ErrExecuteRuntime
	ErrCreateGroup
	ErrCreateUser
	ErrUndefinedGroupEnv
	ErrUndefinedUserEnv
	ErrSetupUser
	ErrLookPath
	ErrSyscallExec
)

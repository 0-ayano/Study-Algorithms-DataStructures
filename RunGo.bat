@setlocal

@cd /d %~dp0
@for %%i in (%*) do (
  Go run %%i
)
ECHO Finish
@ping 127.0.0.1 -n 50 > nul

@endlocal
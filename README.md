# scontrol - Status Control

scontrol is a go package providing go routine control.

## Usage

Use `New()` to get a new controller

```go
sc := scontrol.New()
```

Use `Set()` to set controller into new status; it's okay to set a same status.

```go
sc.Set(scontrol.StatusPause)
```

The go routines being controlled can use `Check()` to get status; it status now is pause, it will block until status changed (means that `Set()` has been called)

```go
go func(){
    for sc.Check() != scontrol.StatusStop {
        // DO its own work
    }
}()
```
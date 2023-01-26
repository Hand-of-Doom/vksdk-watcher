## Installation
Latest version
```shell
go get github.com/Hand-of-Doom/vksdk-watcher/v2@latest
```

### Usage
```go
client := api.NewVK("")

chanErr := make(chan error)
w, startWatching := watcher.NewWatcher(client, chanErr)

w.HandleMessage(`say (?P<phrase>.*)`, func(ctx *watcher.MessageContext, client *api.VK) error {
    phrase := ctx.Params.GetString("phrase")
    _, err := ctx.Answer.Message(phrase).Send()
    return err
})

middleware := func(obj events.WallPostNewObject, client *api.VK) (bool, error) {
    return obj.CreatedBy == 774840378, nil
}

w.WatchWallPostNew(func(obj events.WallPostNewObject, client *api.VK) error {
    p := params.NewMessagesSendBuilder().
        UserID(obj.CreatedBy).
        RandomID(0).
        Message("@id774840378 published a new post")
    _, err := client.MessagesSend(p.Params)
    return err
}, middleware)

w.HandleMessage(`throw error`, func(ctx *watcher.MessageContext, client *api.VK) error {
    return errors.New("something went wrong")
})

go startWatching()

select {
case err := <-chanErr:
    panic(err)
}
```

### Pure VKSDK experience
```go
client := api.NewVK("")

lp, err := longpoll.NewLongPollCommunity(client)
if err != nil {
    panic(err)
}

chanErr := make(chan error)

lp.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
    pattern := regexp.MustCompile(`say (?P<phrase>.*)`)
    if !pattern.MatchString(obj.Message.Text) {
        return
    }
    
    phrasePos := pattern.SubexpIndex("phrase")
    matches := pattern.FindStringSubmatch(obj.Message.Text)
    phrase := matches[phrasePos]
    p := params.NewMessagesSendBuilder().
        Message(phrase).
        RandomID(0).
        PeerID(obj.Message.PeerID)
    _, err = client.MessagesSend(p.Params)
    if err != nil {
        chanErr <- err
    }
})

middleware := func(obj events.WallPostNewObject) (bool, error) {
    return obj.CreatedBy == 774840378, nil
}

lp.WallPostNew(func(ctx context.Context, obj events.WallPostNewObject) {
    ok, err := middleware(obj)
    if err != nil {
        chanErr <- err
    }
    if !ok {
        return
    }
    p := params.NewMessagesSendBuilder().
        UserID(obj.CreatedBy).
        RandomID(0).
        Message("@id774840378 published a new post")
    _, err = client.MessagesSend(p.Params)
    if err != nil {
        chanErr <- err
    }
})

lp.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
    if obj.Message.Text != "throw error" {
        return
    }
    chanErr <- errors.New("something went wrong")
})

go func() {
    // Run() will block executing your program
    err = lp.Run()
    if err != nil {
        chanErr <- err
    }
}()

select {
case err = <-chanErr:
    panic(err)
}
```

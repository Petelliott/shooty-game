package dynconfig


type Dynconfig struct {
    plugindir string
}

func Open(plugindir string) Dynconfig {
    return Dynconfig{plugindir}
}

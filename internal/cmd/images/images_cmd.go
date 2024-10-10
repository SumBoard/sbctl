package images

type Cmd struct {
	Manifest ManifestCmd `cmd:"" help:"Display a manifest of images used by Sumboard and sbctl."`
}

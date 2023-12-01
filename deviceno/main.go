package deviceno

func CustomDeviceNo(prefix, path, listenAddr string) {
	GenerateDeviceNoWithPrefix(prefix, path)
	StartHttp(path, listenAddr, prefix)
}

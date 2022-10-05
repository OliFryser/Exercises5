package client



func main() {
	var opts []grpc.DialOption
	opts = append(
		opts, grpc.WithBlock(), 
		grpc.WithTransportCredentials(insecure.NewCredentials())
	)

	timeContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
}
export default eventHandler(async (event) => {
  return await proxyRequest(event, "http://go:8080/upload");
});

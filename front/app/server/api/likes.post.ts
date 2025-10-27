export default defineEventHandler(async (event) => {
  const idToken = getRequestHeader(event, "authorization");
  const body = await readBody(event);
  const data = await $fetch(`http://go:8080/likes`,
    {
      method: "POST",
      headers: {
        "Authorization": idToken,
        "Content-Type": "application/json",
      },
      body: {
        record_id: Number(body.record_id),
      },
    },
  );
  return data;
});

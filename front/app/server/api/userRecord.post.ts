export default defineEventHandler(async (event) => {
  const idToken = getRequestHeader(event, "authorization");
  const body = await readBody(event);
  const data = await $fetch("http://go:8080/records",
    {
      method: "POST",
      headers: {
        "Authorization": idToken,
        "Content-Type": "application/json",
      },
      body: {
        object_key: body.object_key,
        clean_up_time: Number(body.clean_up_time),
        clean_up_date: new Date(body.clean_up_date),
        comment: body.comment,
      },
    },
  );
  return data;
});

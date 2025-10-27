export default defineEventHandler(async (event) => {
  const userId = getRouterParam(event, "user_id");
  const idToken = getRequestHeader(event, "authorization");
  const data = await $fetch(`http://go:8080/users/${userId}/records`,
    {
      method: "GET",
      headers: {
        "Authorization": idToken,
        "Content-Type": "application/json",
      },
    },
  );
  return data;
});

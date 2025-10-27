export default defineEventHandler(async (event) => {
  const data = $fetch(`http://swagger-api:4010/users/profile`);
  return data;
});

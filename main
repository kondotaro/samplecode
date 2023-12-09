const correctPassword = "your-password-here"; // ここに正しいパスワードを設定してください

export default {
  async fetch(request, env) {
    if (request.method === "POST") {
      const formData = await request.formData();
      const password = formData.get("password");

      if (password === correctPassword) {
        return new Response("認証成功");
      } else {
        return new Response("認証失敗", { status: 401 });
      }
    }

    return new Response(`
      <form method="POST" action="/">
        <label>
          パスワード:
          <input type="password" name="password" required>
        </label>
        <button type="submit">送信</button>
      </form>
    `, {
      headers: { "Content-Type": "text/html" },
    });
  },
};

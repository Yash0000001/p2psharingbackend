package utils

import "fmt"

func ResetPasswordTemplate(username string, link string) string {

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Reset Password</title>
</head>

<body style="background-color:#0f172a;margin:0;padding:0;font-family:Arial,Helvetica,sans-serif;">

<table width="100%%" cellpadding="0" cellspacing="0">
<tr>
<td align="center">

<table width="600" cellpadding="40" cellspacing="0" style="background:#020617;border-radius:12px;margin-top:40px;color:#e2e8f0;">

<tr>
<td align="center">

<h1 style="color:#38bdf8;">Reset Your Password</h1>

<p style="font-size:16px;color:#cbd5f5;">
Hello %s,
</p>

<p style="font-size:15px;color:#cbd5f5;">
We received a request to reset your password.
Click the button below to create a new password.
</p>

<br>

<a href="%s"
style="
background:#38bdf8;
color:#020617;
padding:14px 28px;
text-decoration:none;
font-weight:bold;
border-radius:8px;
display:inline-block;
font-size:16px;">
Reset Password
</a>

<br><br>

<p style="font-size:13px;color:#64748b;">
If you did not request this, you can safely ignore this email.
</p>

<p style="font-size:12px;color:#475569;">
This link will expire in 15 minutes.
</p>

</td>
</tr>

</table>

</td>
</tr>
</table>

</body>
</html>
`, username, link)
}

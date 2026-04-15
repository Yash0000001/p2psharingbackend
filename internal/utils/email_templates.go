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

func WelcomeEmailTemplate(username string, link string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Welcome to Blcak </title>
</head>

<body style="background-color:#000000;margin:0;padding:0;font-family:Arial,Helvetica,sans-serif;">

<table width="100%%" cellpadding="0" cellspacing="0">
<tr>
<td align="center" style="padding: 30px 0;">

<table width="600" cellpadding="40" cellspacing="0" style="background:#0f172a;border-radius:18px;box-shadow:0 24px 80px rgba(15,23,42,.35);color:#e2e8f0;">
<tr>
<td align="center">

<h1 style="margin:0;font-size:32px;color:#38bdf8;">Welcome to P2P Sharing</h1>
<p style="margin:16px 0 24px;font-size:17px;color:#cbd5f5;">Hi %s, thanks for joining our network.</p>

<p style="font-size:15px;line-height:1.8;color:#cbd5f5;max-width:520px;margin:auto;">
We're excited to have you on board. Confirm your email address to finish setting up your account and unlock all sharing features.
</p>

<a href="%s" style="display:inline-block;margin-top:28px;background:#38bdf8;color:#020617;padding:15px 30px;border-radius:12px;text-decoration:none;font-weight:700;font-size:16px;">Verify Your Email</a>

<p style="margin:26px 0 0;font-size:14px;color:#94a3b8;max-width:520px;margin-left:auto;margin-right:auto;line-height:1.6;">
If the button doesn't work, copy and paste the URL below into your browser:
<br><a href="%s" style="color:#38bdf8;word-break:break-all;text-decoration:none;">%s</a>
</p>

<p style="margin-top:28px;font-size:13px;color:#64748b;">Thanks for choosing P2P Sharing.</p>

</td>
</tr>
</table>

</td>
</tr>
</table>

</body>
</html>
`, username, link, link, link)
}

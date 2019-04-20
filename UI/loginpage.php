<html>
<head>
	<?php
            session_start();

            $email="";
            $password="";
            if(isset($_SESSION['loggedin']) && $_SESSION['loggedin']==true){
                header("Location:contact.php");

            }

            if(isset($_POST['email'])&& isset($_POST['password'])){

                if($_POST['email']==$email && $_POST['password']==$password){
                    $_SESSION['loggedin']='true';
                    header("Location:contact.php");
                }
            }
    ?>
</head>
<body>
  <div class="header">
  	<h2><center>Login and grab your favorite burger!</center></h2>
  </div>
	 
  <form method="post" action="trailpage.php">
  
  	<center>
	<table class="input-group">
	<tbody>
	<tr>
		<td><label>E-mail</label></td>
  		<td><input type="email" name="email"></td>
	</tr>
	<br>

	<tr>
  		<td><label>Password</label></td>
  		<td><input type="password" name="password"></td>
	</tr>
	<br>
	<tr>
  	</tbody>
        </table>

	<center><button type="submit" class="btn" name="login_user">Login</button></center>

	<p>
	<center>
  		Not yet a member? <a href="registerpage.php">Sign up</a>
	</center>
	</p>
	</center>
  </form>
</body>
</html>

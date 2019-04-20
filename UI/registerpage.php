<!DOCTYPE html>
<head>
	<div class="header">
	<center><h1>Welcome to the Counter Burger!</h1><center>
	</div>
</head>

<body>
<form method="post" action="registered.php" name="registerform">
	
<center>
  	<table class="input-group">
	<tbody>

	<tr> 
	 <td><label>Name</label></td>
	  <td><input type="text" name="name" value="<?php echo $name; ?>"></td>
	</tr>
	<br>
	
	<tr>		
         <td><label>Address</label></td>
	 <td><input type="text" name="address" value="<?php echo $address; ?>"></td>
	</tr>
	<br>

	<tr>
	 <td><label>Pincode</label></td>
         <td><input type="text" name="pincode" value="<?php echo $pincode; ?>"></td>
	</tr>
	<br>
	
	<tr>
  	  <td><label>Email</label></td>
  	  <td><input type="email" name="email" value="<?php echo $email; ?>"></td>
	</tr>
	<br>

	 <tr>
	<td><label>Password</label></td>
	 <td><input type="password" name="password" value="<?php echo $password; ?>"></td>
	</tr>
	</tbody>
	</table>
	<br>

  	<center><input value="Register" type="submit" class="btn" name="reg_user"></center>

	<p>
	<center>
  		Already a member? <a href="loginpage.php">Sign in</a>
	</center>  
	</p>
	</center>
  </form>
</body>
</html>

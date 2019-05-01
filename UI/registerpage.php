<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<!-- Add icon library -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body {font-family: Mclawsuit  ;}
* {box-sizing: border-box;}

#register_header{color: white  ;
		 background-image: url(burger_texture.jpg);
		 background-position: left top;
		 background-size:300px 300px ; 
		 background-repeat: repeat;
		 padding: 25px;	
		}


.input-container {
  display: -ms-flexbox; /* IE10 */
  display: flex;
  width: 100%;
  margin-bottom: 15px;
}

.icon {
  padding: 10px;
  background: dodgerblue;
  color: white;
  min-width: 50px;
  text-align: center;
}

.input-field {
  width: 100%;
  padding: 10px;
  outline: none;
}

.input-field:focus {
  border: 2px solid dodgerblue;
}

/* Set a style for the submit button */
.btn {
  background-color: dodgerblue;
  color: white;
  padding: 4px 4px 8px 10px;
  border: none;
  cursor: pointer;
  width: 100%;
  opacity: 0.9;
}

.btn:hover {
  opacity: 1;
}

.btn1 {
background-color: dodgerblue;
  color: white;
  padding: 1px 1px 2px 2px;
  text-align: center;
  border: none;
  cursor: pointer;
  width: 100%;
  height: 25%;
  opacity: 0.9;
}

.btn1:hover {
 opacity: 1;
}
</style>
</head>

<body>
<center>
 <div id="register_header" class="header">
	<center><h1>Welcome to the Counter Burger!</h1><center>
	<p></p>
	<p></p>
 </div>
<br>

<form method="post" action="registered.php" name="registerform" style="max-width:500px;margin:auto">

<div>
	<center><h2>Register</h2></center>
</div>

 <div class="input-container">
    <i class="fa fa-user icon"></i>
    <input class="input-field" type="text" value="<?php echo $name; ?>" placeholder="Name" name="nm">
  </div>

<div class="input-container">
	<i class="fa fa-address-card icon"></i>
	<input class="input-field" type="text" value="<?php echo $address; ?>" placeholder="Address" name="address">
</div>

<div class="input-container">
	 <i class="fa fa-address-card-o icon"></i>
	 <input class="input-field" type="text" value="<?php echo $pincode; ?>" pattern="[0-9]{5}" placeholder="Pincode" name="pincode">
</div>

  <div class="input-container">
    <i class="fa fa-envelope icon"></i>
    <input class="input-field" type="email" value="<?php echo $email; ?>" placeholder="Email" name="email">
  </div>
  
  <div class="input-container">
    <i class="fa fa-key icon"></i>
    <input class="input-field" type="password" value="<?php echo $password; ?>" placeholder="Password" name="password">
  </div>

  <button type="submit" class="btn"><h4>Sign Up</h4></button><br>
</form>

<form action="loginpage.php" style="max-width:500px;margin:auto">
<table>
<tbody>
<tr>
<td><h3>Already a member?</h3></td>
<td><button type="submit" class="btn1"><h4>Sign In here!</h4></button></td>
</tr>
</tbody>
</table>
</form>

</center>

</body>
</html>

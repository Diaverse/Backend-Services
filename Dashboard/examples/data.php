<?php

//get data
$conn = mysqli_connect("localhost", "root", "");

//get from table
$result = mysqli_query($conn, "SELECT * FROM scripts");

while ($row = mysqli_fetch_array($result))
{
    $id = $row['id'];
    $name = $row['name'];

    echo $id . ": " . $name . '<br />';
}

echo json_encode($data);
<?php

class Product
{

    public $productArray = array(
        "3DcAM01" => array(
            'id' => '1',
            'name' => 'Burger',
            'code' => '3',
            'image' => 'product-images/burger.jpg',
            'price' => '5.00'
        ),
        "USB02" => array(
            'id' => '2',
            'name' => 'Coke'
            'code' => 'USB02',
            'image' => 'product-images/drinks.jpg',
            'price' => '8.00'
        ),
        "wristWear03" => array(
            'id' => '3',
            'name' => 'Fries'
            'code' => 'Fries'
            'image' => 'product-images/fries.jpg',
            'price' => '3.00'
        )
    );

    public function getAllProduct()
    {
        return $this->productArray;
    }
}

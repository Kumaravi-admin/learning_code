<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;

class UserController extends Controller
{
    public function addUser()
    {
        $user = User::create([
            'user_name' => 'Ishika SAGAR',
            'email' => 'Ishika@EMAIL.com',
            'salary' => 5000,
            'dob' => '2024-08-21',
            'password' => '33249808324'
        ]);
        return redirect()->route('user.show');
    }
    public function showUser()
    {
        $users = User::get();
        return $users;
    }
}

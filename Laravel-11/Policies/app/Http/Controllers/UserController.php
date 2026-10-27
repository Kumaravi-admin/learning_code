<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class UserController extends Controller
{
    public function loginPage()
    {
        if (Auth::check()) {
            return redirect()->route('books.index');
        }
        return view('login');
    }

    public function login(Request $Request)
    {
        $credential = $Request->validate([
            'email' => 'required|email',
            'password' => 'required'
        ]);
        if (Auth::attempt($credential)) {
            return redirect()->route('books.index');
        }
        return redirect()->back();
    }

    public function logout()
    {
        Auth::logout();
        return redirect()->route('home');
    }


}

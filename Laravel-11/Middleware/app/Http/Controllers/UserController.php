<?php

namespace App\Http\Controllers;

use App\Models\User;
use Auth;
use Illuminate\Http\Request;

class UserController extends Controller
{
    public function homePage()
    {
        if (Auth::guest()) {
            return view('welcome');
        } else {
            return redirect()->route('dashboard');
        }
    }

    public function loginPage()
    {
        return view('login');
    }
    public function registerPage()
    {
        return view('register');
    }

    public function register(Request $request)
    {
        $data = $request->validate([
            'name' => 'required',
            'email' => 'required|email',
            'userage' => 'required|integer',
            'role' => 'required',
            'password' => 'required|confirmed'
        ]);
        $user = User::create($data);

        if ($user) {

            return redirect()->route('loginPage');
        }
    }

    public function login(Request $request)
    {
        $credentials = $request->validate([
            'email' => 'required|email',
            'password' => 'required'
        ]);
        if (Auth::attempt($credentials)) {
            if (Auth::check()) {
                return redirect()->route('dashboard');
            } else {
                return redirect()->route('loginPage');
            }
        } else {
            return redirect()->route('loginPage');
        }
    }

    public function dashboardPage()
    {
        return view('dashboard');
    }
    public function innerPage()
    {
        return view('inner');
    }

    public function logout()
    {
        Auth::logout();
        return redirect()->route('loginPage');
    }
}

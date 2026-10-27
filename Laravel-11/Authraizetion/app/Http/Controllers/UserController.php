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

    public function register(Request $request)
    {
        $data = $request->validate([
            'name' => 'required',
            'email' => 'required|email',
            'password' => 'required|confirmed'
        ]);
        $user = User::create($data);

        if ($user) {

            return redirect()->route('login');
        }
    }

    public function login(Request $request)
    {
        $creadentials = $request->validate([
            'email' => 'required|email',
            'password' => 'required'
        ]);
        if (Auth::attempt($creadentials)) {
            return redirect()->route('dashboard');
        }
    }

    public function dashboardPage()
    {
        if (Auth::check()) {
            return view('dashboard');
        } else {
            return redirect()->route('login');
        }
    }
    public function innerPage()
    {
        if (Auth::check()) {
            return view('inner');
        } else {
            return redirect()->route('login');
        }
    }

    public function logout()
    {
        Auth::logout();
        return redirect()->route('login');
    }
}

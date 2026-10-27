<?php

namespace App\Http\Controllers;

use App\Models\Post;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Gate;

class UserController extends Controller
{
    public function loginPage()
    {
        return view('login');
    }

    public function loginMatch(Request $request)
    {
        $credential = $request->validate([
            'email' => 'required|email',
            'password' => 'required',
        ]);

        if (Auth::attempt($credential)) {
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
        // Gate::authorize('isAdmin');
        return view('dashboard');

        // if (Gate::allows('isAdmin')) {
        //     return view('dashboard');
        // } else {
        //     return 'Access deinei';
        // }

    }

    public function profilePage(int $userID)
    {


        // if (Gate::allows('view-profile', $userID)) {
        //     $user = User::findOrFail($userID);
        //     return $user;
        // } else {
        //     return 'Wrong user id';
        // }

        Gate::authorize('view-profile', $userID);
        $user = User::findOrFail($userID);
        return $user;
    }

    public function showPost()
    {
        $posts = Post::where('user_id', Auth::id())->get();
        return view('post', compact('posts'));
    }

    public function postUpdate($postId)
    {
        $post = Post::find($postId);

        $userId = $post->user_id;

        Gate::authorize('update-post', $userId);

        $post = Post::findOrFail($postId);
        return $post;
    }

    public function logout()
    {
        Auth::logout();
        return redirect()->route('loginPage');
    }
}
// Gate::any(['view-profile'])
// Gate::none(['view-profile'])
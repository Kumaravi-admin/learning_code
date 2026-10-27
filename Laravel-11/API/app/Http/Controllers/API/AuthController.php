<?php

namespace App\Http\Controllers\API;

use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Validator;
use App\Http\Controllers\Fascia;

class AuthController extends Controller
{
    /*
    |--------------------------------------------------------------------------
    | Register the new user : signup() method
    |--------------------------------------------------------------------------
    |
    | 1 define validation rule for request data
    | 2 using Fascia::validate method to validate request data : validate($reqest data, validation rule)
    | 3 if fascia::validate() method return false then return error respose using fasica:: error() method
    | 4 stor new user data in User table
    | 5 checking if user save then return success responce if any error return error responce
    |
    */

    public function signup(Request $request)
    {
        $validation_rules = [
            'name' => 'required',
            'email' => 'required|email|unique:users,email',
            'password' => 'required',
        ];

        $validateUser = Fascia::validate($request, $validation_rules);

        if ($validateUser !== true) {
            return Fascia::error('Error While Validating Data', $validateUser);
        }

        $user = User::create([
            'name' => $request->name,
            'email' => $request->email,
            'password' => $request->password,
        ]);

        if ($user) {
            return Fascia::success($user, 'User Register Successfully');
        } else {
            return Fascia::error('Something Wrong User data not save. Please try again', $user);
        }

    }







    /*
    |--------------------------------------------------------------------------
    | Authenticate the user : login() method
    |--------------------------------------------------------------------------
    | 
    | 1) Define validation rule for request data
    | 2) Validate data using fascia::validate() method
    | 3) if fascia::validate() method return false then return error respose using fasica:: error() method
    | 4) Authenticate user using Auth::attempt(email id, Password) method, if user 
    |    Authentication success then return token in respose if Authentication false
    |    return error using fascia::error(error message) method
    |    
    */

    public function login(Request $request)
    {
        $validation_rules = [
            'email' => 'required',
            'password' => 'required',
        ];

        $validateUser = Fascia::validate($request, $validation_rules);

        if ($validateUser !== true) {
            return Fascia::error('Error While Validating Data', $validateUser, 404);
        }

        if (Auth::attempt(['email' => $request->email, 'password' => $request->password])) {
            return response()->json([
                'status' => true,
                'message' => 'User Login Successfully.',
                'token' => Auth::user()->createToken('API Token')->plainTextToken,
                'token_type' => 'bearer'
            ], 200);

        } else {
            return Fascia::error('Email & Password dose not match.');
        }
    }



    /*
    |--------------------------------------------------------------------------
    | Logout the user : logout() method
    |--------------------------------------------------------------------------
    |
    | 1) Make user varibale, using $user->tokens->delete to delete user all tokens
    | 2) Regenrate session token
    | 3) Return success responce using fascia::usccess() method
    |
    */

    public function logout(Request $request)
    {
        $user = $request->user();
        $user->tokens()->delete();
        session()->regenerate();
        // $user->tokens(API KEY NAME)->delete();
        return Fascia::success($user, 'Logout Successfully');

    }



}

<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class SessionController extends Controller
{
    public function show_session()
    {
        $value = session()->get('name');
        echo $value;
        echo session('status');

        // $value = session('lang');
        // echo $value;

        // $value = session('name', 'hello');
        // echo $value;

        // $value = session()->all();
        // echo '<pre>';
        // print_r($value);
        // echo '</pre>';

        // $value = session()->except(['name']);
        // echo '<pre>';
        // print_r($value);
        // echo '</pre>';


        // $value = session()->only(['name']);
        // echo '<pre>';
        // print_r($value);
        // echo '</pre>';

        // if (session()->has('name')) {
        //     $value = session()->get('name');
        //     echo $value;
        // } else {
        //     echo 'name key dosnot exists';
        // }

        // if (session()->exists('name')) {
        //     $value = session()->get('name');
        //     echo $value;
        // } else {
        //     echo 'name key dosnot exists';
        // }


    }
    public function store_session()
    {
        // session(['name' => 'Yahoo Baba']);
        // session()->put('name', 'Yahoo Baba');
        session([
            'name' => 'Yahoo Baba',
            'lang' => 'laravel'
        ]);
        session()->flash('status', 'session saved successfully.');
        return redirect('/session/show');


        // session()->increment('count', $incrimentBy = 2);
        // session()->decrement('count', $decrimentby = 2);
        // session()->regenerate();

    }

    public function delete_session()
    {
        // session()->forget('name');
        // session()->forget(['name', 'lang']);
        // session()->flush();
        session()->invalidate();
    }
}

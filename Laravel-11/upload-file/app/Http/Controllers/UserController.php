<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;

class UserController extends Controller
{
    public function index()
    {
        $users = User::get();
        return view('upload-file', compact('users'));
    }

    public function store(Request $request)
    {
        // $file = $request->file('photo');

        // $request->validate([
        //     'photo' => 'required|mimes:png,jpeg,jpg|max:5000|dimensions:min_width:10,min_height:10px'
        // ]);

        // $path = $request->file('photo')->store('images', 'public');
        // file(key)->stor(folder name,inside Public or Local folder)                                   
        // return $path;


        // $file = $request->file('photo');
        // $request->validate(['photo' => 'required|mimes:png,jpeg,jpg|max:5000']);
        // $fileName = time() . '_' . $file->getClientOriginalName();
        // $path = $request->file('photo')->storeAs('images', $fileName, 'public');
        // return $path;


        // $file = $request->file('photo');
        // // $extention = $file->getClientOriginalExtension();
        // $extention = $file->extension();
        // return $extention;

        // $file = $request->file('photo');
        // $path = $file->hashName();
        // return $path;

        // $file = $request->file('photo');
        // $type = $file->getClientMimeType();
        // return $type;

        // $file = $request->file('photo');
        // $file_size = $file->getSize();
        // return $file_size;

        // $file = $request->file('photo');
        // $fileName = time() . '_' . $file->hashName();
        // $file->move(public_path('uploads'), $fileName);

        // User::create([
        //     'file_name' => $fileName
        // ]);

        // return redirect()->route('user.index');

        $file = $request->file('photo');
        $fileName = time() . '_' . $file->hashName();
        $file->storeAs('images', $fileName, 'public');

        User::create([
            'file_name' => $fileName
        ]);

        return redirect()->route('user.index');

    }
    public function edit(string $id)
    {
        $user = User::find($id);
        return view('update-file', compact('user'));
    }
    public function update(Request $request, string $id)
    {
        $user = User::find($id);

        if ($request->hasFile('photo')) {

            $image_path = public_path('storage/images/') . $user->file_name;
            if (file_exists($image_path)) {
                @unlink($image_path);
            }

            $file = $request->file('photo');
            $fileName = time() . '_' . $file->hashName();
            $file->storeAs('images', $fileName, 'public');

            $user->whereId($id)->update([
                'file_name' => $fileName
            ]);
            return redirect()->route('user.index');
        }
    }
    public function destroy(string $id)
    {
        $user = User::find($id);
        $user->delete();

        $image_path = public_path('storage/images/') . $user->file_name;

        if (file_exists($image_path)) {
            @unlink($image_path);
            return redirect()->route('user.index');
        }
        return 'file not exists';

    }
}

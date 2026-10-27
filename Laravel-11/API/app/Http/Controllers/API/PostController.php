<?php

namespace App\Http\Controllers\API;

use App\Models\Post;
use Illuminate\Http\Request;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Validator;
use App\Http\Controllers\Fascia;

class PostController extends Controller
{
    /**
     * Display a listing of the resource.
     */

    public function index()
    {
        $data['posts'] = Post::all();

        return response()->json([
            'status' => true,
            'message' => 'All Post Data',
            'data' => $data,
        ], 200);

        // return Fascia::success($data, 'All Post Data');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        $validatePost = Validator::make(
            $request->all(),
            [
                'title' => 'required',
                'description' => 'required',
                'image' => 'required|mimes:png,jpg,jpeg,gif',
            ]
        );
        // $validation_rules = [
        //     'title' => 'required',
        //     'description' => 'required',
        //     'image' => 'required|mimes:png,jpg,jpeg,gif',
        // ];

        // $validate = Fascia::validate($request, $validation_rules);

        // if ($validate !== true) {
        //     return Fascia::error('Validation Error in Fascia', $validate, 400);
        // }


        if ($validatePost->fails()) {
            return response()->json([
                'status' => false,
                'message' => 'Validation Error',
                'errors' => $validatePost->errors()->all()
            ], 401);
        }

        $img = $request->image;
        $imageName = time() . '.' . $img->getClientOriginalExtension();
        $img->move(public_path() . '/uploads', $imageName);



        // $imageName = Fascia::uploadFile($request->file('image'), '/uploads');

        // if ($imageName[0] === false) {
        //     return Fascia::error('File not upload', $imageName['error']);
        // }

        $post = Post::create([
            'title' => $request->title,
            'description' => $request->description,
            'image' => $imageName,
        ]);


        return response()->json([
            'status' => true,
            'message' => 'New Post Created Successfully.',
            'post' => $post
        ], 200);

        // return Fascia::success($post, 'New Post Created Successfully');

    }

    /**
     * Display the specified resource.
     */
    public function show(string $id)
    {
        $data['post'] = Post::findOrFail($id);


        return response()->json([
            'status' => true,
            'message' => 'Your Single Post.',
            'data' => $data
        ], 200);


    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        $validatePost = Validator::make(
            $request->all(),
            [
                'title' => 'required',
                'description' => 'required',
                'image' => 'required|mimes:png,jpg,jpeg,gif',
            ]
        );

        if ($validatePost->fails()) {
            return response()->json([
                'status' => false,
                'message' => 'Validation Error',
                'errors' => $validatePost->errors()->all()
            ], 401);
        }

        $postImage = Post::where('id', $id)->get('image');

        // $imageName = Fascia::updateFile('', $postImage[0]->image, '/uploads');


        if ($request->file('image') != '') {

            $path = public_path() . '/uploads';
            if ($postImage[0]->image != '' && $postImage[0]->image != null) {
                $old_path = $path . $postImage[0]->image;
                if (file_exists($old_path)) {
                    unlink($old_path);
                }
            }
            $img = $request->image;
            $imageName = time() . '.' . $img->getClientOriginalExtension();
            $img->move(public_path() . '/uploads', $imageName);

        } else {
            $imageName = $postImage->image;
        }


        $post = Post::where(['id' => $id])->update([
            'title' => $request->title,
            'description' => $request->description,
            'image' => $imageName,
        ]);

        return response()->json([
            'status' => true,
            'message' => 'Post Update Successfully.',
            'post' => $post
        ], 200);
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        $imagePath = Post::select('image')->where('id', $id)->get();

        $filePath = public_path() . '/uploads/' . $imagePath[0]['image'];
        if (file_exists($filePath)) {
            unlink($filePath);
        }

        $post = Post::where('id', $id)->delete();

        return response()->json([
            'status' => true,
            'message' => 'Post Delete Successfully.',
            'post' => $post
        ], 200);
    }
}

<?php

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;
use Illuminate\Validation\ValidationException;

class Fascia extends Controller
{

    public static function updateFile($file, $old_file_name, $folder, )
    {

        if ($file !== '') {

            $path = public_path() . $folder . '/';

            if ($old_file_name !== '' && $old_file_name !== null) {

                $old_path = $path . $old_file_name;
                if (file_exists($old_path)) {
                    unlink($old_path);
                }
            }

            $new_file_name = time() . '.' . $file->getClientOriginalExtension();

            if ($file->move($path, $new_file_name)) {
                return $new_file_name;
            } else {
                return [0 => false, 'error' => 'File not uploaded successfully'];
            }
        } else {
            return $old_file_name;
        }



    }





    public static function uploadFile($file, $folder)
    {
        $fileName = time() . '.' . $file->getClientOriginalExtension();
        $file->move(public_path() . $folder, $fileName);

        if (file_exists(public_path() . $folder . '/' . $fileName)) {
            return $fileName;
        } else {
            return [0 => false, 'error' => 'File not upload something wrong while uploading please try again'];
        }

    }
    public static function validate($data, $validation_rules)
    {
        $validator = Validator::make($data->all(), $validation_rules);
        if ($validator->fails()) {
            
            return $validator->errors()->all();
            
        }
        return true;
    }

    public static function success($result, $message)
    {
        return response()->json([
            'success' => true,
            'message' => $message,
            'data' => $result,
        ], 200);
    }

    public static function error($errormessage, $error = [], $code = 401)
    {
        $responce = [
            'success' => false,
            'message' => $errormessage,
        ];
        if (!empty($error)) {
            $responce['error'] = $error;
        }
        return response()->json($responce, $code);
    }



}

//rule_start
            'email' => ['required',  'string', ],
            'company' => ['required',  'string', ],
            'extra' => [ 'string', ],
            'position' => ['required',  'string', ],
            'tel' => ['required',  'string', function ($attribute, $value, $fail) {
                if (preg_match("/^1[3456789]\d{9}$/", $value)) {
                    return true;
                } else {
                    $fail('手机号格式错误');
                }
            },
],
            'name' => ['required',  'string', ],
            'company' => ['required',  'string', ],
            'extra' => [ 'string', ],
            'position' => ['required',  'string', ],
            'tel' => ['required',  'string', function ($attribute, $value, $fail) {
                if (preg_match("/^1[3456789]\d{9}$/", $value)) {
                    return true;
                } else {
                    $fail('手机号格式错误');
                }
            },
],
            'name' => ['required',  'string', ],
            'email' => ['required',  'string', ],
            'tel' => ['required',  'string', function ($attribute, $value, $fail) {
                if (preg_match("/^1[3456789]\d{9}$/", $value)) {
                    return true;
                } else {
                    $fail('手机号格式错误');
                }
            },
],
            'name' => ['required',  'string', ],
            'email' => ['required',  'string', ],
            'company' => ['required',  'string', ],
            'extra' => [ 'string', ],
            'position' => ['required',  'string', ],
            'name' => ['required',  'string', ],
            'email' => ['required',  'string', ],
            'company' => ['required',  'string', ],
            'extra' => [ 'string', ],
            'position' => ['required',  'string', ],
            'tel' => ['required',  'string', function ($attribute, $value, $fail) {
                if (preg_match("/^1[3456789]\d{9}$/", $value)) {
                    return true;
                } else {
                    $fail('手机号格式错误');
                }
            },
],
            'tel' => ['required',  'string', function ($attribute, $value, $fail) {
                if (preg_match("/^1[3456789]\d{9}$/", $value)) {
                    return true;
                } else {
                    $fail('手机号格式错误');
                }
            },
],
            'name' => ['required',  'string', ],
            'email' => ['required',  'string', ],
            'company' => ['required',  'string', ],
            'extra' => [ 'string', ],
            'position' => ['required',  'string', ],
            'position' => ['required',  'string', ],
            'tel' => ['required',  'string', function ($attribute, $value, $fail) {
                if (preg_match("/^1[3456789]\d{9}$/", $value)) {
                    return true;
                } else {
                    $fail('手机号格式错误');
                }
            },
],
            'name' => ['required',  'string', ],
            'email' => ['required',  'string', ],
            'company' => ['required',  'string', ],
            'extra' => [ 'string', ],
            'email' => ['required',  'string', ],
            'company' => ['required',  'string', 'max:10', 'min:1',],
            'extra' => [ 'string', ],
            'position' => ['required',  'string', ],
            'tel' => ['required',  'string', function ($attribute, $value, $fail) {
                if (preg_match("/^1[3456789]\d{9}$/", $value)) {
                    return true;
                } else {
                    $fail('手机号格式错误');
                }
            },
],
            'name' => ['required',  'string', ],
            'company' => ['required',  'string', 'max:10', 'min:1',],
            'extra' => [ 'string', ],
            'position' => ['required',  'string', ],
            'tel' => ['required',  'string', function ($attribute, $value, $fail) {
                if (preg_match("/^1[3456789]\d{9}$/", $value)) {
                    return true;
                } else {
                    $fail('手机号格式错误');
                }
            },
],
            'name' => ['required',  'string', 'max:10', ],
            'email' => ['required',  'string', ],
//rule_end

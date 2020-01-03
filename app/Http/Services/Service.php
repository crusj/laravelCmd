<?php
//service_start
    //服务详情
    public function show(int $id) :array {

        return [];
    }
    //服务评论
    public function comments(int $id) :array {
        //
        $page = request()->get('page', '');
        //
        $perPage = request()->get('per_page', '');
        return [];
    }
    //评论服务
    public function storeComment(int $id) :int {

        return 0;
    }
    //预约咨询
    public function storeOrder(int $id) :int {

        return 0;
    }
    //收藏
    public function storeCollect(int $id) :int {

        return 0;
    }
//service_end

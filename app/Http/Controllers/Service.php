<?php
namespace App\Http\Controllers;

use App\Http\Controllers\Base;

class Service extends Base
{
//controller_start
    //服务详情
    public function show($id) {
        $service = ServiceFactory::service();
        $data = $service->show($id);
        if(empty($data)) {
            $this->failObject('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //服务评论
    public function comments($id) {
        $service = ServiceFactory::service();
        $data = $service->comments($id);
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //评论服务
    public function storeComment($id) {
        $service = ServiceFactory::service();
        $id = $service->storeComment($id);
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //预约咨询
    public function storeOrder($id) {
        $service = ServiceFactory::service();
        $id = $service->storeOrder($id);
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //收藏
    public function storeCollect($id) {
        $service = ServiceFactory::service();
        $id = $service->storeCollect($id);
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
//controller_end
}

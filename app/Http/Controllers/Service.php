<?php
namespace App\Http\Controllers;

use App\Http\Controllers\Base;

class Service extends Base
{
//controller_start
    //服务分类
    public function index() {
        $service = ServiceFactory::service();
        $data = $service->index();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //所有服务
    public function index() {
        $service = ServiceFactory::service();
        $data = $service->index();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //服务详情
    public function index(int $id) {
        $service = ServiceFactory::service();
        $data = $service->index(int $id);
        if(empty($data)) {
            $this->failObject('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //服务评论
    public function comment(int $id) {
        $service = ServiceFactory::service();
        $data = $service->comment(int $id);
        if(empty($data)) {
            $this->failObject('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //评论服务
    public function storeComment(int $id) {
        $service = ServiceFactory::service();
        $id = $service->storeComment(int $id);
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //预约咨询
    public function storeOrder(int $id) {
        $service = ServiceFactory::service();
        $id = $service->storeOrder(int $id);
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //收藏
    public function storeCollect(int $id) {
        $service = ServiceFactory::service();
        $i = $service->storeCollect(int $id);
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
//controller_end
}

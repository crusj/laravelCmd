<?php
//controller_start
    //授权登录
    public function storeLogin() {
        $service = ServiceFactory::user();
        $id = $service->storeLogin();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //获取验证码
    public function getCode() {
        $service = ServiceFactory::user();
        $data = $service->getCode();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //更新手机号
    public function storeUpdateTel() {
        $service = ServiceFactory::user();
        $id = $service->storeUpdateTel();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //验证手机
    public function updateVerifyCode() {
        $service = ServiceFactory::user();
        $ok = $service->updateVerifyCode();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //文件上传
    public function storeUploadFile() {
        $service = ServiceFactory::user();
        $id = $service->storeUploadFile();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //反馈与帮助
    public function storeFeedback() {
        $service = ServiceFactory::user();
        $id = $service->storeFeedback();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //隐私政策
    public function privacy() {
        $service = ServiceFactory::user();
        $data = $service->privacy();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //用户协议
    public function agreement() {
        $service = ServiceFactory::user();
        $data = $service->agreement();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //信息更改
    public function updateUpdateInfo() {
        $service = ServiceFactory::user();
        $ok = $service->updateUpdateInfo();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //个人信息
    public function info() {
        $service = ServiceFactory::user();
        $data = $service->info();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //授权登录
    public function storeLogin() {
        $service = ServiceFactory::user();
        $id = $service->storeLogin();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //获取验证码
    public function getCode() {
        $service = ServiceFactory::user();
        $data = $service->getCode();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //更新手机号
    public function storeUpdateTel() {
        $service = ServiceFactory::user();
        $id = $service->storeUpdateTel();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //验证手机
    public function updateVerifyCode() {
        $service = ServiceFactory::user();
        $ok = $service->updateVerifyCode();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //文件上传
    public function storeUploadFile() {
        $service = ServiceFactory::user();
        $id = $service->storeUploadFile();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //反馈与帮助
    public function storeFeedback() {
        $service = ServiceFactory::user();
        $id = $service->storeFeedback();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //隐私政策
    public function privacy() {
        $service = ServiceFactory::user();
        $data = $service->privacy();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //用户协议
    public function agreement() {
        $service = ServiceFactory::user();
        $data = $service->agreement();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //信息更改
    public function updateUpdateInfo() {
        $service = ServiceFactory::user();
        $ok = $service->updateUpdateInfo();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //个人信息
    public function info() {
        $service = ServiceFactory::user();
        $data = $service->info();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //授权登录
    public function storeLogin() {
        $service = ServiceFactory::user();
        $id = $service->storeLogin();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //获取验证码
    public function getCode() {
        $service = ServiceFactory::user();
        $data = $service->getCode();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //更新手机号
    public function storeUpdateTel() {
        $service = ServiceFactory::user();
        $id = $service->storeUpdateTel();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //验证手机
    public function updateVerifyCode() {
        $service = ServiceFactory::user();
        $ok = $service->updateVerifyCode();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //文件上传
    public function storeUploadFile() {
        $service = ServiceFactory::user();
        $id = $service->storeUploadFile();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //反馈与帮助
    public function storeFeedback() {
        $service = ServiceFactory::user();
        $id = $service->storeFeedback();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //隐私政策
    public function privacy() {
        $service = ServiceFactory::user();
        $data = $service->privacy();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //用户协议
    public function agreement() {
        $service = ServiceFactory::user();
        $data = $service->agreement();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //信息更改
    public function updateUpdateInfo() {
        $service = ServiceFactory::user();
        $ok = $service->updateUpdateInfo();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //个人信息
    public function info() {
        $service = ServiceFactory::user();
        $data = $service->info();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //授权登录
    public function storeLogin() {
        $service = ServiceFactory::user();
        $id = $service->storeLogin();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //获取验证码
    public function getCode() {
        $service = ServiceFactory::user();
        $data = $service->getCode();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //更新手机号
    public function storeUpdateTel() {
        $service = ServiceFactory::user();
        $id = $service->storeUpdateTel();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //验证手机
    public function updateVerifyCode() {
        $service = ServiceFactory::user();
        $ok = $service->updateVerifyCode();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //文件上传
    public function storeUploadFile() {
        $service = ServiceFactory::user();
        $id = $service->storeUploadFile();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //反馈与帮助
    public function storeFeedback() {
        $service = ServiceFactory::user();
        $id = $service->storeFeedback();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //隐私政策
    public function privacy() {
        $service = ServiceFactory::user();
        $data = $service->privacy();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //用户协议
    public function agreement() {
        $service = ServiceFactory::user();
        $data = $service->agreement();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //信息更改
    public function updateUpdateInfo() {
        $service = ServiceFactory::user();
        $ok = $service->updateUpdateInfo();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //个人信息
    public function info() {
        $service = ServiceFactory::user();
        $data = $service->info();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //授权登录
    public function storeLogin() {
        $service = ServiceFactory::user();
        $id = $service->storeLogin();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //获取验证码
    public function getCode() {
        $service = ServiceFactory::user();
        $data = $service->getCode();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //更新手机号
    public function storeUpdateTel() {
        $service = ServiceFactory::user();
        $id = $service->storeUpdateTel();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //验证手机
    public function updateVerifyCode() {
        $service = ServiceFactory::user();
        $ok = $service->updateVerifyCode();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //文件上传
    public function storeUploadFile() {
        $service = ServiceFactory::user();
        $id = $service->storeUploadFile();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //反馈与帮助
    public function storeFeedback() {
        $service = ServiceFactory::user();
        $id = $service->storeFeedback();
        if($id <= 0) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->success(['id' => $id]);
        }
    }
    //隐私政策
    public function privacy() {
        $service = ServiceFactory::user();
        $data = $service->privacy();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //用户协议
    public function agreement() {
        $service = ServiceFactory::user();
        $data = $service->agreement();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //信息更改
    public function updateUpdateInfo() {
        $service = ServiceFactory::user();
        $ok = $service->updateUpdateInfo();
        if($ok === false) {
            $this->failObject($service->getErrorMsg());
        }else {
            $this->successNull();
        }
    }
    //个人信息
    public function info() {
        $service = ServiceFactory::user();
        $data = $service->info();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
//controller_end

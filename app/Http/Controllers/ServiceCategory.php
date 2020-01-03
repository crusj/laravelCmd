//controller_start
    //服务分类
    public function index() {
        $service = ServiceFactory::serviceCategory();
        $data = $service->index();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //所有服务
    public function allService() {
        $service = ServiceFactory::serviceCategory();
        $data = $service->allService();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //服务分类
    public function index() {
        $service = ServiceFactory::serviceCategory();
        $data = $service->index();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //所有服务
    public function allService() {
        $service = ServiceFactory::serviceCategory();
        $data = $service->allService();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //服务分类
    public function index() {
        $service = ServiceFactory::serviceCategory();
        $data = $service->index();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
    //所有服务
    public function allService() {
        $service = ServiceFactory::serviceCategory();
        $data = $service->allService();
        if(empty($data)) {
            $this->failArray('暂无数据');
        }else {
            $this->success($data);
        }
    }
//controller_end

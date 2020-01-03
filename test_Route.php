<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\User;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:api')->get('/user', function (Request $request) {
    return $request->user();
});
//不强校验token
Route::middleware(['jmhc.check.token:false'])->group(function () {
    Route::prefix("user")->group(function () {
        Route::post("login", 'User@login');
    });
    Route::get("notice", "Notice@index");
    Route::get("notice/{id}", "Notice@show");
    Route::get('banner', 'Banner@index');

    Route::prefix('service')->group(function () {
        Route::get('category', 'ServiceCategory@index');
    });

});
//强校验token
Route::middleware(['jmhc.check.token'])->group(function () {
    Route::prefix("user")->group(function () {
        Route::post('updateTel', 'User@updateTel');
    });
//route_start
    //产品服务
    Route::prefix('service_collect')->group(function(){
        //收藏记录
        Route::get('/','ServiceCollect@index');
    });
    //活动
    Route::prefix('activity')->group(function(){
        //收藏
        Route::post('/{id}/collect','Activity@storeCollect');
        //活动列表
        Route::get('/','Activity@index');
        //活动详情
        Route::get('/{id}','Activity@show');
        //立即报名
        Route::post('/{id}/join','Activity@storeJoin');
        //预约咨询
        Route::post('/{id}/order','Activity@storeOrder');
    });
    //Banner
    Route::prefix('banner')->group(function(){
        //列表
        Route::get('/','Banner@index');
    });
    //平台公告
    Route::prefix('notice')->group(function(){
        //列表
        Route::get('/','Notice@index');
        //详情
        Route::get('/{id}','Notice@show');
    });
    //产品服务
    Route::prefix('service_category')->group(function(){
        //服务分类
        Route::get('/','ServiceCategory@index');
        //所有服务
        Route::get('/allService','ServiceCategory@allService');
    });
    //产品服务
    Route::prefix('service')->group(function(){
        //服务详情
        Route::get('/{id}','Service@show');
        //服务评论
        Route::get('/{id}/comments','Service@comments');
        //评论服务
        Route::post('/{id}/comment','Service@storeComment');
        //预约咨询
        Route::post('/{id}/order','Service@storeOrder');
        //收藏
        Route::post('/{id}/collect','Service@storeCollect');
    });
    //产品服务
    Route::prefix('service_order')->group(function(){
        //预约记录
        Route::get('/','ServiceOrder@index');
    });
    //新闻
    Route::prefix('news')->group(function(){
        //分类和标签
        Route::get('/tagsAndTypes','News@tagsAndTypes');
        //新闻列表
        Route::get('/','News@index');
        //新闻详情
        Route::get('/{id}','News@show');
        //新闻评论
        Route::get('/{id}/comments','News@comments');
        //评论新闻
        Route::post('/{id}/comment','News@storeComment');
        //预约咨询
        Route::post('/{id}/order','News@storeOrder');
        //收藏
        Route::post('/{id}/collect','News@storeCollect');
        //点赞
        Route::post('/{id}/thumb','News@storeThumb');
    });
    //新闻视频
    Route::prefix('video')->group(function(){
        //视频列表
        Route::get('/','Video@index');
        //视频详情
        Route::get('/{id}','Video@show');
        //视频评论
        Route::get('/{id}/comments','Video@comments');
        //评论视频
        Route::post('/{id}/comment','Video@storeComment');
        //点赞
        Route::post('/{id}/thumb','Video@storeThumb');
    });
    //中介服务
    Route::prefix('middle')->group(function(){
        //中介列表
        Route::get('/','Middle@index');
        //中介详情
        Route::get('/{id}','Middle@show');
        //申请联系方式
        Route::post('/{id}/applyTel','Middle@storeApplyTel');
        //评论中介
        Route::post('/{id}/comment','Middle@storeComment');
        //预约咨询
        Route::post('/{id}/order','Middle@storeOrder');
        //收藏
        Route::post('/{id}/collect','Middle@storeCollect');
        //点赞
        Route::post('/{id}/thumb','Middle@storeThumb');
    });
    //个人中心
    Route::prefix('user')->group(function(){
        //授权登录
        Route::post('/login','User@login');
        //获取验证码
        Route::get('/getCode','User@getCode');
        //更新手机号
        Route::post('/updateTel','User@updateTel');
        //验证手机
        Route::put('/verifyCode','User@verifyCode');
        //文件上传
        Route::post('/uploadFile','User@uploadFile');
        //反馈与帮助
        Route::post('/feedback','User@feedback');
        //隐私政策
        Route::get('/privacy','User@privacy');
        //用户协议
        Route::get('/agreement','User@agreement');
        //信息更改
        Route::put('/updateInfo','User@updateInfo');
        //个人信息
        Route::get('/info','User@info');
    });
//route_end
//admin_route_start
//admin_route_end
});


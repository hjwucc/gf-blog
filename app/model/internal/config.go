// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal



// Config is the golang structure for table config.
type Config struct {
    Id          int    `orm:"id,primary"   json:"id"`           //                              
    YpuBucket   string `orm:"ypu_bucket"   json:"ypu_bucket"`   // 又拍云存储服务名             
    YpuOperator string `orm:"ypu_operator" json:"ypu_operator"` // 又拍云操作员                 
    YpyPassword string `orm:"ypy_password" json:"ypy_password"` // 又拍云密码                   
    YpyHosts    string `orm:"ypy_hosts"    json:"ypy_hosts"`    // 又拍云自定义 Hosts 映射关系  
}
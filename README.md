# identity
基于fabric1.4学生身份及学历管理系统

链码层                                 业务层                             Web层
addIde                                SaveIde                            AddIde
queryIdeByCertNoAndName               FindIdeByCertNoAndName             FindCertByNoAndName
queryIdeInfoByIDNumber                FindIdeInfoByIDNumber              FindByID
updateIdeByCertBo                     ModifyIde                          Modify

web层调用业务层，业务层调用链码层，实现浏览器客户端对账本的操作

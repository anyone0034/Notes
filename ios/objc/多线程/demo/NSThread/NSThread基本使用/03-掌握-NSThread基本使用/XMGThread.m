//
//  XMGThread.m
//  03-掌握-NSThread基本使用
//
//  Created by xiaomage on 16/2/18.
//  Copyright © 2016年 小码哥. All rights reserved.
//

#import "XMGThread.h"

@implementation XMGThread

-(void)dealloc
{
    NSLog(@"dealloc----%@",[NSThread currentThread]);
}
@end

//
//  XMGViewController.m
//  09-通过Xib加载控制器(掌握)
//
//  Created by xiaomage on 16/1/14.
//  Copyright © 2016年 小码哥. All rights reserved.
//

#import "XMGViewController.h"

@interface XMGViewController ()

@end

@implementation XMGViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view.
}


-(instancetype)initWithNibName:(NSString *)nibNameOrNil bundle:(NSBundle *)nibBundleOrNil{
    NSLog(@"%s",__func__);
    return [super initWithNibName:nibNameOrNil bundle:nibBundleOrNil];
}


- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

/*
#pragma mark - Navigation

// In a storyboard-based application, you will often want to do a little preparation before navigation
- (void)prepareForSegue:(UIStoryboardSegue *)segue sender:(id)sender {
    // Get the new view controller using [segue destinationViewController].
    // Pass the selected object to the new view controller.
}
*/

@end

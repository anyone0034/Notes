//
//  XMGStatusCell.h
//  备课07-自定义不等高的cell-frame
//
//  Created by FTD_ZHC on 15/9/22.
//  Copyright (c) 2015年 小码哥. All rights reserved.
//

#import <UIKit/UIKit.h>

@class XMGStatusFrame;
@interface XMGStatusCell : UITableViewCell

/**
 *  frame模型数据
 */
@property (nonatomic ,strong)XMGStatusFrame * statusFrame;

@end

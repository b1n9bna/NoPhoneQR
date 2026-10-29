#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Generate an enhanced high quality comparison GIF for NoPhoneQR README with live benchmark metrics.
"""
import math
import os
from PIL import Image, ImageDraw, ImageFont

W, H = 940, 580
FONT_PATH = "/System/Library/Fonts/PingFang.ttc"
FONT_ENG = "/System/Library/Fonts/Supplemental/Arial.ttf"

def get_font(size, bold=False):
    index = 1 if bold else 0
    try:
        return ImageFont.truetype(FONT_PATH, size, index=index)
    except:
        try:
            return ImageFont.truetype(FONT_ENG, size)
        except:
            return ImageFont.load_default()

font_title = get_font(24, bold=True)
font_subtitle = get_font(13)
font_header = get_font(16, bold=True)
font_body = get_font(13)
font_small = get_font(11)
font_tag = get_font(12, bold=True)
font_stat = get_font(11, bold=True)

FRAMES_COUNT = 32

def draw_round_rect(draw, box, radius, fill, outline=None, width=1):
    x1, y1, x2, y2 = box
    draw.rounded_rectangle([x1, y1, x2, y2], radius=radius, fill=fill, outline=outline, width=width)

def render_frame(t):
    img = Image.new("RGBA", (W, H), (13, 17, 23, 255)) # Dark GitHub background #0D1117
    draw = ImageDraw.Draw(img)

    # 1. Top Header Banner
    draw_round_rect(draw, [20, 16, W - 20, 84], 12, fill=(22, 27, 34, 255), outline=(48, 54, 61, 255), width=1)
    
    # Title
    draw.text((40, 26), "📱❌ NoPhoneQR: 电脑屏幕扫码终极对决", fill=(240, 246, 252), font=font_title)
    draw.text((40, 58), "告别举手机扫电脑屏幕 ｜ 实测内存仅 9.1MB · 体积 3.3MB · 零 Electron", fill=(139, 148, 158), font=font_subtitle)

    # Dynamic Status Pill on Top Right
    draw_round_rect(draw, [W - 235, 30, W - 40, 70], 18, fill=(35, 134, 54, 40), outline=(46, 160, 67, 180), width=1)
    if t < 0.5:
        draw.text((W - 215, 41), "⚡ 效率提升 4000% (0.8s)", fill=(63, 185, 80), font=font_tag)
    else:
        draw.text((W - 218, 41), "🚀 峰值内存仅 9.1MB (极省)", fill=(63, 185, 80), font=font_tag)

    # 2. Left Column: Traditional Approach (Red Theme)
    left_x1, left_y1, left_x2, left_y2 = 20, 100, W//2 - 10, H - 20
    draw_round_rect(draw, [left_x1, left_y1, left_x2, left_y2], 12, fill=(22, 27, 34, 255), outline=(248, 81, 73, 120), width=1)

    # Left Header
    draw_round_rect(draw, [left_x1 + 16, left_y1 + 16, left_x2 - 16, left_y1 + 54], 8, fill=(248, 81, 73, 30), outline=(248, 81, 73, 100))
    draw.text((left_x1 + 28, left_y1 + 24), "❌ 传统方式 (Traditional)  ·  7 个步骤", fill=(248, 81, 73), font=font_header)

    trad_steps = [
        ("1. 摸口袋掏出手机 📱", 0.05),
        ("2. 面容或指纹解锁 🔓", 0.18),
        ("3. 桌面翻找并打开微信/相机 🔍", 0.32),
        ("4. 身体后倾、手举手机对准屏幕对焦 📷", 0.50),
        ("5. 手机端加载满是追踪参数的网页 ⏳", 0.68),
        ("6. 复制链接，发送给微信“文件传输助手” 📤", 0.82),
        ("7. 电脑端微信接收并在浏览器打开 🏁", 0.95),
    ]

    sy = left_y1 + 70
    for idx, (step_text, trigger_t) in enumerate(trad_steps):
        is_done = t >= trigger_t
        is_current = (idx == 0 and t < trigger_t) or (idx > 0 and t >= trad_steps[idx-1][1] and t < trigger_t)

        box_color = (33, 38, 45, 255) if not is_current else (48, 54, 61, 255)
        text_color = (201, 209, 217) if is_done else ((240, 246, 252) if is_current else (110, 118, 129))
        dot_color = (248, 81, 73) if is_done else ((210, 153, 34) if is_current else (48, 54, 61))

        draw_round_rect(draw, [left_x1 + 16, sy, left_x2 - 16, sy + 36], 6, fill=box_color)
        draw.ellipse([left_x1 + 26, sy + 13, left_x1 + 36, sy + 23], fill=dot_color)
        draw.text((left_x1 + 46, sy + 9), step_text, fill=text_color, font=font_body)
        sy += 42

    # Left Bottom Summary Bar
    left_progress = min(1.0, t / 0.95)
    draw_round_rect(draw, [left_x1 + 16, left_y2 - 62, left_x2 - 16, left_y2 - 16], 8, fill=(33, 38, 45, 255))
    draw.text((left_x1 + 24, left_y2 - 50), f"耗时进度: {int(left_progress * 35)} 秒", fill=(139, 148, 158), font=font_small)
    # Progress track
    track_x = left_x1 + 130
    draw_round_rect(draw, [track_x, left_y2 - 46, left_x2 - 24, left_y2 - 38], 4, fill=(48, 54, 61))
    draw_round_rect(draw, [track_x, left_y2 - 46, track_x + int((left_x2 - left_x1 - 154) * left_progress), left_y2 - 38], 4, fill=(248, 81, 73))
    draw.text((left_x1 + 24, left_y2 - 32), "状态: 动作繁琐 · 打断心流 · 效率低下", fill=(248, 81, 73), font=font_small)

    # 3. Right Column: NoPhoneQR (Emerald Theme)
    right_x1, right_y1, right_x2, right_y2 = W//2 + 10, 100, W - 20, H - 20
    draw_round_rect(draw, [right_x1, right_y1, right_x2, right_y2], 12, fill=(19, 35, 27, 255), outline=(46, 160, 67, 180), width=2)

    # Right Header
    draw_round_rect(draw, [right_x1 + 16, right_y1 + 16, right_x2 - 16, right_y1 + 54], 8, fill=(35, 134, 54, 40), outline=(46, 160, 67, 120))
    draw.text((right_x1 + 28, right_y1 + 24), "✅ 使用 NoPhoneQR  ·  1 步即达", fill=(63, 185, 80), font=font_header)

    # Spec Badges Strip
    spec_w = (right_x2 - right_x1 - 48) // 3
    specs = [
        ("🧠 内存 9.1MB", (35, 134, 54, 50), (63, 185, 80)),
        ("📦 体积 3.3MB", (56, 139, 253, 50), (88, 166, 255)),
        ("⚡ 启动 <50ms", (210, 153, 34, 50), (227, 179, 65)),
    ]
    for si, (stxt, sbg, sfg) in enumerate(specs):
        bx1 = right_x1 + 16 + si * (spec_w + 8)
        draw_round_rect(draw, [bx1, right_y1 + 62, bx1 + spec_w, right_y1 + 88], 6, fill=sbg, outline=sfg)
        draw.text((bx1 + 10, right_y1 + 68), stxt, fill=sfg, font=font_stat)

    # Interactive Stage Simulation in Right Column
    stage_box = [right_x1 + 16, right_y1 + 96, right_x2 - 16, right_y2 - 128]
    draw_round_rect(draw, stage_box, 8, fill=(13, 17, 23, 255), outline=(48, 54, 61, 255))

    # Fake Screen Window inside stage
    draw.text((stage_box[0] + 16, stage_box[1] + 12), "快捷键触发: Cmd + Shift + Q (或 Alt + Q)", fill=(88, 166, 255), font=font_body)

    # Simulated QR box
    qr_cx, qr_cy = (stage_box[0] + stage_box[2]) // 2, (stage_box[1] + stage_box[3]) // 2 + 10
    qr_w, qr_h = 96, 96
    qr_x1, qr_y1 = qr_cx - qr_w//2, qr_cy - qr_h//2
    qr_x2, qr_y2 = qr_cx + qr_w//2, qr_cy + qr_h//2

    # Draw simulated QR pattern
    draw.rectangle([qr_x1, qr_y1, qr_x2, qr_y2], fill=(240, 246, 252))
    # Corner markers
    for mx, my in [(qr_x1 + 6, qr_y1 + 6), (qr_x2 - 26, qr_y1 + 6), (qr_x1 + 6, qr_y2 - 26)]:
        draw.rectangle([mx, my, mx + 20, my + 20], fill=(0, 0, 0))
        draw.rectangle([mx + 4, my + 4, mx + 16, my + 16], fill=(240, 246, 252))
        draw.rectangle([mx + 8, my + 8, mx + 12, my + 12], fill=(0, 0, 0))
    # Inner random-like grid
    for gx in range(qr_x1 + 34, qr_x2 - 10, 12):
        for gy in range(qr_y1 + 12, qr_y2 - 10, 12):
            if (gx + gy) % 24 == 0:
                draw.rectangle([gx, gy, gx + 8, gy + 8], fill=(0, 0, 0))

    # Crosshair and Selection Box animation on NoPhoneQR
    if t < 0.25:
        ratio = t / 0.25
        cur_x = qr_x1 - 30 + int(30 * ratio)
        cur_y = qr_y1 - 30 + int(30 * ratio)
        draw.line([cur_x - 10, cur_y, cur_x + 10, cur_y], fill=(63, 185, 80), width=2)
        draw.line([cur_x, cur_y - 10, cur_x, cur_y + 10], fill=(63, 185, 80), width=2)
    elif t < 0.60:
        draw.rectangle([qr_x1 - 8, qr_y1 - 8, qr_x2 + 8, qr_y2 + 8], outline=(63, 185, 80), width=2)
        scan_progress = ((t - 0.25) / 0.35)
        scan_y = qr_y1 - 8 + int((qr_h + 16) * scan_progress)
        draw.line([qr_x1 - 8, scan_y, qr_x2 + 8, scan_y], fill=(0, 255, 136), width=3)
    else:
        draw.rectangle([qr_x1 - 8, qr_y1 - 8, qr_x2 + 8, qr_y2 + 8], outline=(63, 185, 80), width=2)
        draw_round_rect(draw, [qr_cx - 86, qr_cy - 18, qr_cx + 86, qr_cy + 18], 8, fill=(35, 134, 54, 235), outline=(63, 185, 80))
        draw.text((qr_cx - 72, qr_cy - 9), "⚡ 0.8s 秒级解析完成", fill=(255, 255, 255), font=font_tag)

    # Pop-up Notification Card
    notif_y = right_y2 - 116
    draw_round_rect(draw, [right_x1 + 16, notif_y, right_x2 - 16, right_y2 - 16], 8, fill=(22, 27, 34, 255), outline=(46, 160, 67, 120))
    if t >= 0.50:
        draw.text((right_x1 + 28, notif_y + 12), "🌐 系统默认浏览器已在新标签页弹开！", fill=(63, 185, 80), font=font_body)
        draw.text((right_x1 + 28, notif_y + 36), "🛡️ 已自动剔除 6 个跟踪参数，净化真实短链", fill=(139, 148, 158), font=font_small)
        draw.text((right_x1 + 28, notif_y + 58), "🔗 https://github.com/b1n9bna/NoPhoneQR", fill=(88, 166, 255), font=font_small)
        draw.text((right_x1 + 28, notif_y + 78), "✅ 内存已释放归零 (Run & Done 无驻留)", fill=(63, 185, 80), font=font_small)
    else:
        draw.text((right_x1 + 28, notif_y + 24), "⏳ 正在等待快捷框选...", fill=(139, 148, 158), font=font_body)
        draw.text((right_x1 + 28, notif_y + 54), "💡 占用内存: 0 MB (未触发时不消耗任何硬件资源)", fill=(139, 148, 158), font=font_small)

    return img.convert("RGB")

frames = []
for i in range(FRAMES_COUNT):
    progress = i / float(FRAMES_COUNT - 1)
    frames.append(render_frame(progress))

# Add 12 pause frames at the end so users can inspect the result
for _ in range(12):
    frames.append(render_frame(1.0))

out_path = "/Users/colin569/社会生产劳动/GitHub/NoPhoneQR/assets/demo.gif"
frames[0].save(
    out_path,
    save_all=True,
    append_images=frames[1:],
    duration=110,
    loop=0,
    optimize=True
)
print("Updated GIF saved successfully to:", out_path)

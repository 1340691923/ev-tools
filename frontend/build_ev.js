import fs from 'fs/promises';
import path from 'path';
import url from 'url';
const __filename = url.fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// 更新文件路径，指向 dist 文件夹下的 index.html
const filePath = path.resolve(__dirname, 'dist', 'index.html');

async function updateHtml() {
    try {

        // 读取文件内容
        const content = await fs.readFile(filePath, 'utf8');

        // 使用正则表达式替换内容
        const modifiedContent = content.replaceAll("'/__dynamic_base__", "window.__dynamic_base__+'");

        // 将修改后的内容写回同一个文件
        await fs.writeFile(filePath, modifiedContent, 'utf8');
        console.log('dist/index.html 文件内容已更新。');
    } catch (err) {
        console.error('处理文件时出错:', err);
    }
}
await updateHtml()

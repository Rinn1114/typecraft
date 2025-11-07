import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import App from './App';

// 我们将模拟全局的 fetch 函数
beforeEach(() => {
  global.fetch = jest.fn();
});

test('fetches and displays a message from the backend', async () => {
  // 准备一个模拟的成功响应
  const mockMessage = 'Hello from a mock API!';
  const mockResponse = { message: mockMessage };

  // 配置我们的假 fetch 函数
  (global.fetch as jest.Mock).mockResolvedValue({
    ok: true,
    json: jest.fn().mockResolvedValue(mockResponse),
  });

  // 渲染 App 组件
  render(<App />);

  // 初始状态：检查是否显示了加载信息
  expect(screen.getByText(/Loading message from backend.../i)).toBeInTheDocument();

  // 等待异步操作完成，并检查最终的后端消息是否已显示
  await waitFor(() => {
    expect(screen.getByText(mockMessage, { selector: 'strong' })).toBeInTheDocument();
  });

  // 确认加载信息已经消失
  expect(screen.queryByText(/Loading message from backend.../i)).not.toBeInTheDocument();
});

test('displays an error message if the fetch fails', async () => {
  // 配置我们的假 fetch 函数以返回一个错误
  (global.fetch as jest.Mock).mockRejectedValue(new Error('Network error'));

  // 渲染 App 组件
  render(<App />);

  // 等待异步操作完成，并检查错误消息是否已显示
  await waitFor(() => {
    expect(screen.getByText(/Error: Network error/i)).toBeInTheDocument();
  });
});

package geoip

import (
	"fmt"
	"net"
	"project/zj"
)

type TrieNode struct {
	sub   [2]*TrieNode
	isEnd bool // 标记是否为 IP 段的结束
}

type IPTrie struct {
	root *TrieNode
}

var theTrie = NewIPTrie()

func NewIPTrie() *IPTrie {
	return &IPTrie{root: &TrieNode{}}
}

// 插入 IP 段
func (t *IPTrie) Insert(cidr string) error {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("invalid CIDR: %v", err)
	}

	node := t.root
	for _, bit := range cidrToBits(ipNet) {
		if node.sub[bit] == nil {
			node.sub[bit] = &TrieNode{}
		}
		node = node.sub[bit]
	}
	node.isEnd = true
	return nil
}

// 检查 IP 是否在任意段中
func (t *IPTrie) Contains(ip net.IP) bool {
	node := t.root
	for _, bit := range ipToBits(ip) {
		if node == nil {
			return false
		}
		node = node.sub[bit]
		if node != nil && node.isEnd {
			return true // 提前终止，匹配成功
		}
	}
	return false
}

// 将 IPNet 转换为比特数组
func cidrToBits(ipNet *net.IPNet) []int {
	var bits []int
	ip := ipNet.IP
	mask, _ := ipNet.Mask.Size()
	for i := 0; i < mask; i++ {
		bits = append(bits, int((ip[i/8]>>(7-i%8))&1))
	}
	return bits
}

// 将 IP 转换为比特数组
func ipToBits(ip net.IP) []int {
	ip = ip.To4() // 支持 IPv4
	var bits []int
	for i := 0; i < 32; i++ {
		bits = append(bits, int((ip[i/8]>>(7-i%8))&1))
	}
	return bits
}

func init() {
	// 构建 Trie
	for _, cidr := range ipRanges {
		if err := theTrie.Insert(cidr); err != nil {
			zj.W("Failed to insert CIDR %s: %v\n", cidr, err)
		}
	}
}

package menu

type NodeSon struct {
	NodeSonText    string
	NodeSonPriv    string
	NodeSonPrivAct string
}

type Node struct {
	NodeText    string
	NodeUrl     string
	NodePriv    string
	NodePrivAct string
	PrivChild   []NodeSon
}

type Menu struct {
	MenuText    string
	MenuPriv    string
	MenuPrivAct string
	MenuIcon    string
	Nodes       []Node
}

var MenuList []Menu

func GetMenu() []Menu {
	MenuList = []Menu{
		{
			MenuText:    "设置",
			MenuPriv:    "setting",
			MenuPrivAct: "get",
			MenuIcon:    "mdi mdi-settings",
			Nodes: []Node{
				{
					NodeText:    "管理员管理",
					NodeUrl:     "/admin/setting/adminuser/index",
					NodePriv:    "/admin/setting/adminuser/index",
					NodePrivAct: "get",
					PrivChild: []NodeSon{
						{
							NodeSonText:    "添加管理员",
							NodeSonPriv:    "/admin/setting/adminuser/add",
							NodeSonPrivAct: "get",
						},
						{
							NodeSonText:    "编辑管理员",
							NodeSonPriv:    "/admin/setting/adminuser/edit",
							NodeSonPrivAct: "get",
						},
						{
							NodeSonText:    "保存管理员",
							NodeSonPriv:    "/admin/setting/adminuser/save",
							NodeSonPrivAct: "post",
						},
					},
				},
				{
					NodeText:    "角色管理",
					NodeUrl:     "/admin/setting/admingroup/index",
					NodePriv:    "/admin/setting/admingroup/index",
					NodePrivAct: "get",
					PrivChild: []NodeSon{
						{
							NodeSonText:    "添加角色",
							NodeSonPriv:    "/admin/setting/admingroup/add",
							NodeSonPrivAct: "get",
						},
						{
							NodeSonText:    "编辑角色",
							NodeSonPriv:    "/admin/setting/admingroup/edit",
							NodeSonPrivAct: "get",
						},
						{
							NodeSonText:    "保存角色",
							NodeSonPriv:    "/admin/setting/admingroup/save",
							NodeSonPrivAct: "post",
						},
					},
				},
				{
					NodeText:    "系统日志",
					NodeUrl:     "/admin/setting/system/index",
					NodePriv:    "/admin/setting/system/index",
					NodePrivAct: "get",
					PrivChild: []NodeSon{
						{
							NodeSonText:    "获取目录",
							NodeSonPriv:    "/admin/setting/system/getdir",
							NodeSonPrivAct: "get",
						},
						{
							NodeSonText:    "读取日志",
							NodeSonPriv:    "/admin/setting/system/view",
							NodeSonPrivAct: "get",
						},
					},
				},
			},
		},
		{
			MenuText:    "示例",
			MenuPriv:    "demo",
			MenuPrivAct: "get",
			MenuIcon:    "mdi mdi-format-list-bulleted",
			Nodes: []Node{
				{
					NodeText:    "附件管理",
					NodeUrl:     "/admin/demo/show",
					NodePriv:    "/admin/demo/show",
					NodePrivAct: "get",
					PrivChild: []NodeSon{
						{
							NodeSonText:    "上传文件",
							NodeSonPriv:    "/admin/demo/upload",
							NodeSonPrivAct: "post",
						},
					},
				},
			},
		},
	}
	return MenuList
}

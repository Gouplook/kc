 #!/usr/bin/python
import redis

notes_key = "rpc_card_note:notes"
requirements_key = "rpc_card_note:requirements"

notes = (\
"此卡只限本店使用",\
"使用本卡时请注意门店的营业时间",\
"使用本卡时请注意卡项的次数","消费者确认消费一次后，此卡可在平台上进行美转让",\
"此卡不能转让",\
"此卡属实名制卡",\
"此卡为虚拟卡不能报失，请妥善保管好您的用户账号", "如您有什么疑问，请向本店咨询",\
"此卡解释权归本公司所有")


requirements = (\
"提前一周预约",\
"提前1天预约",\
"提前2天预约",\
"提前3天预约",\
"不需要预约")

host = "192.168.1.246"
port=6379

def main():
	global notes_key, requirements_key
	global notes, requirements
	global host, port
	conn = redis.Redis(host, port)
	conn.delete(notes_key)
	conn.delete(requirements_key)
	conn.lpush(notes_key, *notes)
	conn.lpush(requirements_key, *requirements )
	conn.close()
	return

if (__name__ == "__main__"):
	main()
<style class='fox_global_style'>
	div.fox_html_content { line-height: 1.5; }
</style>
<div>
	<table border='1' bordercolor='#000000' cellpadding='0' cellspacing='0' style='font-size: 10pt; border-collapse:collapse; border:none' width='80%'>
		<caption>
			<font size='2' face='Verdana'>
				贴吧签到结果
			</font>
		</caption>
		<tbody>
			<tr>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							ID
						</div>
					</font>
				</td>
				<td width='15%' style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							等级
						</div>
					</font>
				</td>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							名称
						</div>
					</font>
				</td>
				<td width='10%' style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							结果
						</div>
					</font>
				</td>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							内容
						</div>
					</font>
				</td>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							签到排名
						</div>
					</font>
				</td>
			</tr>
            {{range .Tiebas}}
			<tr>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							{{.Id}}
						</div>
					</font>
				</td>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							等级：{{.LevelId}}级&nbsp;&nbsp;称号：{{.LevelName}}
							<br/>
							经验：
							<span style='color: #f15a23'>
								{{.CurScore}}
							</span>
							/{{.LevelupScore}}
						</div>
					</font>
				</td>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							{{.Name}}
							<br/>
							{{.Slogan}}
						</div>
					</font>
				</td>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							<span style='color: green'>
                            {{.AddScore}}
							</span>
						</div>
					</font>
				</td>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							{{.ErrorCode}}
							<br/>
							{{.ErrorMsg}}
						</div>
					</font>
				</td>
				<td style='border: solid 1 #000000' nowrap=''>
					<font size='2' face='Verdana'>
						<div>
							今日本吧第
							<span style='color: #f15a23'>
								{{.Rank}}
							</span>
							个签到
							<br/>
							连续签到：
							<span style='color: #f15a23'>
								{{.SignKeep}}
							</span>
							天,累计签到：
							<span style='color: #f15a23'>
								{{.SignTotal}}
							</span>
							天
						</div>
					</font>
				</td>
			</tr>
            {{end}}
			<tr>
				<td style='text-align: center;' colspan='6'>
					本次签到共：
					<span style='color: #f15a23'>
						{{.Count}}
					</span>
					个贴吧；签到累计获得
					<span style='color: #f15a23'>
                        {{.TotalScore}}
					</span>
					经验值
				</td>
			</tr>
            
		</tbody>
	</table>
	<span id='_FoxCURSOR'>
	</span>
</div>
<hr id='FMSigSeperator' style='width: 210px; height: 1px;' color='#b5c4df' size='1' align='left'>
<div>
	<span id='_FoxFROMNAME'>
		<div style='MARGIN: 10px; FONT-FAMILY: verdana; FONT-SIZE: 10pt'>
			<div>
			</div>
			<div>
				<u>
					TEL：15000846364
				</u>
			</div>
			<span style='color: rgb(0, 0, 0); background-color: rgba(0, 0, 0, 0);'>
			</span>
		</div>
	</span>
</div>